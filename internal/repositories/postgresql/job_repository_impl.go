package postgresql

import (
	"context"
	"database/sql"
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/helpers"
	"dz-jobs-api/internal/models"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type SQLJobRepository struct {
	db *sql.DB
}

func NewJobRepository(db *sql.DB) repositoryInterfaces.JobRepository {
	return &SQLJobRepository{
		db: db,
	}
}

func (r *SQLJobRepository) CreateJob(ctx context.Context, job *models.Job) error {
	query := `
        INSERT INTO jobs (
            title, description, location, salary_range, required_skills, recruiter_id, created_at, updated_at, status, job_type
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
        ) RETURNING job_id
    `

	err := r.db.QueryRow(
		query,
		job.Title, job.Description, job.Location, job.SalaryRange, job.RequiredSkills, job.RecruiterID,
		job.CreatedAt, job.UpdatedAt, job.Status, job.JobType,
	).Scan(&job.ID)

	if err != nil {
		return fmt.Errorf("repository: failed to create job: %w", err)
	}
	return nil
}

func (r *SQLJobRepository) GetJobDetails(ctx context.Context, jobID int64, recruiterID uuid.UUID) (*models.Job, error) {
	if err := r.ValidateJobOwnership(ctx, jobID, recruiterID); err != nil {
		return nil, err
	}

	query := `SELECT job_id, title, description, location, salary_range, required_skills, recruiter_id, created_at, updated_at, status, job_type
              FROM jobs WHERE job_id = $1`

	job := &models.Job{}
	err := r.db.QueryRow(query, jobID).Scan(
		&job.ID, &job.Title, &job.Description, &job.Location, &job.SalaryRange, &job.RequiredSkills,
		&job.RecruiterID, &job.CreatedAt, &job.UpdatedAt, &job.Status, &job.JobType,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("repository: job not found")
		}
		return nil, fmt.Errorf("repository: failed to fetch job by ID: %w", err)
	}
	return job, nil
}

func (r *SQLJobRepository) GetJobListingsByStatus(ctx context.Context, status string, recruiterID uuid.UUID) ([]*models.Job, error) {

	query := `SELECT job_id, title, description, location, salary_range, required_skills, recruiter_id, created_at, updated_at, status
              FROM jobs WHERE status = $1 AND recruiter_id = $2`

	rows, err := r.db.Query(query, status, recruiterID)
	if err != nil {
		return nil, fmt.Errorf("repository: failed to fetch jobs by status: %w", err)
	}
	defer rows.Close()

	var jobs []*models.Job
	for rows.Next() {
		job := &models.Job{}
		err := rows.Scan(
			&job.ID, &job.Title, &job.Description, &job.Location, &job.SalaryRange, &job.RequiredSkills,
			&job.RecruiterID, &job.CreatedAt, &job.UpdatedAt, &job.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("repository: failed to scan job: %w", err)
		}
		if job.RecruiterID != recruiterID {
			return nil, fmt.Errorf("repository: unauthorized access, recruiter does not own job with ID %d", job.ID)
		}
		jobs = append(jobs, job)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("repository: rows error: %w", err)
	}

	return jobs, nil
}

func (r *SQLJobRepository) UpdateJob(ctx context.Context, jobID int64, recruiterID uuid.UUID, job *models.Job) error {
	query := `UPDATE jobs SET 
        title = $1, description = $2, location = $3, salary_range = $4, required_skills = $5, 
        recruiter_id = $6, updated_at = $7, status = $8, job_type = $9
        WHERE job_id = $10`

	result, err := r.db.Exec(
		query,
		job.Title, job.Description, job.Location, job.SalaryRange, job.RequiredSkills,
		job.RecruiterID, job.UpdatedAt, job.Status, job.JobType, jobID,
	)

	if err != nil {
		return fmt.Errorf("repository: failed to update job: %w", err)
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		return fmt.Errorf("repository: failed to check rows affected: %w", err)
	} else if rowsAffected == 0 {
		return errors.New("repository: job not found")
	}

	return nil
}

func (r *SQLJobRepository) DeactivateJob(ctx context.Context, jobID int64, recruiterID uuid.UUID) error {

	query := `UPDATE jobs SET 
        status = $1, 
        updated_at = $2 
        WHERE job_id = $3`

	result, err := r.db.Exec(query, "closed", time.Now(), jobID)
	if err != nil {
		return fmt.Errorf("repository: failed to update job status to closed: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository: failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("repository: job not found")
	}

	return nil
}

func (r *SQLJobRepository) RepostJob(ctx context.Context, jobID int64, recruiterID uuid.UUID) error {

	query := `UPDATE jobs SET 
        status = $1, 
        updated_at = $2 
        WHERE job_id = $3`

	result, err := r.db.Exec(query, "open", time.Now(), jobID)
	if err != nil {
		return fmt.Errorf("repository: failed to update job status to open: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository: failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("repository: job not found")
	}

	return nil
}

func (r *SQLJobRepository) DeleteJob(ctx context.Context, jobID int64, recruiterID uuid.UUID) error {

	deleteQuery := "DELETE FROM jobs WHERE job_id = $1"
	result, err := r.db.Exec(deleteQuery, jobID)
	if err != nil {
		return fmt.Errorf("repository: failed to delete job: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository: failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("repository: job not found")
	}

	return nil
}

func (r *SQLJobRepository) ValidateJobOwnership(ctx context.Context, jobID int64, recruiterID uuid.UUID) error {
	query := `SELECT recruiter_id FROM jobs WHERE job_id = $1`
	row := r.db.QueryRow(query, jobID)

	var ownerID uuid.UUID
	if err := row.Scan(&ownerID); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("repository: job not found: %w", err)
		}
		return fmt.Errorf("repository: failed to check job ownership: %w", err)
	}

	if ownerID != recruiterID {
		return fmt.Errorf("repository: unauthorized access, recruiter does not own the job")
	}

	return nil
}

func (r *SQLJobRepository) GetAllJobs(ctx context.Context) ([]*models.Job, error) {
	query := `SELECT job_id, title, description, location, salary_range, required_skills, recruiter_id, created_at, updated_at, status, job_type
              FROM jobs`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("repository: failed to fetch all jobs: %w", err)
	}
	defer rows.Close()

	var jobs []*models.Job
	for rows.Next() {
		job := &models.Job{}
		err := rows.Scan(
			&job.ID, &job.Title, &job.Description, &job.Location, &job.SalaryRange, &job.RequiredSkills,
			&job.RecruiterID, &job.CreatedAt, &job.UpdatedAt, &job.Status, &job.JobType,
		)
		if err != nil {
			return nil, fmt.Errorf("repository: failed to scan job: %w", err)
		}
		jobs = append(jobs, job)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("repository: rows error: %w", err)
	}

	return jobs, nil
}

func (r *SQLJobRepository) GetJobListings(ctx context.Context, filters request.JobFilters) ([]*models.Job, error) {
	query := `SELECT job_id, title, description, location, salary_range, required_skills, recruiter_id, created_at, updated_at, status, job_type
              FROM jobs WHERE 1=1`

	args := []interface{}{}
	paramCount := 1

	if filters.Status != "" {
		query += fmt.Sprintf(" AND status = $%d", paramCount)
		args = append(args, filters.Status)
		paramCount++
	}

	if filters.JobType != "" {
		query += fmt.Sprintf(" AND job_type = $%d", paramCount)
		args = append(args, filters.JobType)
		paramCount++
	}

	if filters.Location != "" {
		query += fmt.Sprintf(" AND location ILIKE $%d", paramCount)
		args = append(args, "%"+filters.Location+"%")
		paramCount++
	}

	if filters.SalaryRangeMin > 0 || filters.SalaryRangeMax > 0 {
		salaryRangeQuery, salaryRangeArgs := helpers.ConvertSalaryRange(filters.SalaryRangeMin, filters.SalaryRangeMax)
		query += fmt.Sprintf(" AND (%s)", salaryRangeQuery)
		args = append(args, salaryRangeArgs...)
		paramCount += len(salaryRangeArgs)
	}

	if len(filters.RequiredSkills) > 0 {
		skillConditions := []string{}
		for _, skill := range filters.RequiredSkills {
			skillConditions = append(skillConditions, fmt.Sprintf("required_skills ILIKE $%d", paramCount))
			args = append(args, "%"+skill+"%")
			paramCount++
		}
		query += " AND (" + strings.Join(skillConditions, " OR ") + ")"
	}

	if filters.Keyword != "" {
		query += fmt.Sprintf(" AND (title ILIKE $%d OR description ILIKE $%d)",
			paramCount, paramCount+1)
		searchTerm := "%" + filters.Keyword + "%"
		args = append(args, searchTerm, searchTerm)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("repository: failed to fetch jobs with filters: %w", err)
	}
	defer rows.Close()

	var jobs []*models.Job
	for rows.Next() {
		job := &models.Job{}
		err := rows.Scan(
			&job.ID, &job.Title, &job.Description, &job.Location, &job.SalaryRange, &job.RequiredSkills,
			&job.RecruiterID, &job.CreatedAt, &job.UpdatedAt, &job.Status, &job.JobType,
		)
		if err != nil {
			return nil, fmt.Errorf("repository: failed to scan job: %w", err)
		}
		jobs = append(jobs, job)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("repository: rows error: %w", err)
	}

	return jobs, nil
}

func (r *SQLJobRepository) GetJobDetailsPublic(ctx context.Context, jobID int64) (*models.Job, error) {
	query := `SELECT job_id, title, description, location, salary_range, required_skills, recruiter_id, created_at, updated_at, status, job_type
              FROM jobs WHERE job_id = $1`

	row := r.db.QueryRow(query, jobID)
	job := &models.Job{}
	err := row.Scan(
		&job.ID, &job.Title, &job.Description, &job.Location, &job.SalaryRange, &job.RequiredSkills,
		&job.RecruiterID, &job.CreatedAt, &job.UpdatedAt, &job.Status, &job.JobType,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("repository: failed to fetch job by ID: %w", err)
	}
	return job, nil
}
