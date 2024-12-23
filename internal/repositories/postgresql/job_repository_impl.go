package postgresql

import (
	"database/sql"
	"dz-jobs-api/internal/models"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces"
	"errors"
	"fmt"
	"time"
)

type SQLJobRepository struct {
	db *sql.DB
}

func NewJobRepository(db *sql.DB) repositoryInterfaces.JobRepository {
	return &SQLJobRepository{
		db: db,
	}
}

func (r *SQLJobRepository) CreateJob(job *models.Job) error {
	query := `
        INSERT INTO jobs (
            title, description, location, salary_range, required_skills, recruiter_id, created_at, updated_at, status
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9
        ) RETURNING job_id
    `

	var jobID int64
	err := r.db.QueryRow(
		query,
		job.Title, job.Description, job.Location, job.SalaryRange, job.RequiredSkills, job.RecruiterID,
		job.CreatedAt, job.UpdatedAt, job.Status,
	).Scan(&jobID)

	if err != nil {
		return fmt.Errorf("repository: failed to create job: %w", err)
	}

	job.JobID = jobID
	return nil
}

func (r *SQLJobRepository) GetJobDetails(jobID int64) (*models.Job, error) {
	query := `SELECT job_id, title, description, location, salary_range, required_skills, recruiter_id, created_at, updated_at, status
              FROM jobs WHERE job_id = $1`

	row := r.db.QueryRow(query, jobID)
	job := &models.Job{}
	err := row.Scan(
		&job.JobID, &job.Title, &job.Description, &job.Location, &job.SalaryRange, &job.RequiredSkills,
		&job.RecruiterID, &job.CreatedAt, &job.UpdatedAt, &job.Status,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("repository: failed to fetch job by ID: %w", err)
	}
	return job, nil
}

func (r *SQLJobRepository) GetJobListingsByStatus(status string) ([]*models.Job, error) {
	query := `SELECT job_id, title, description, location, salary_range, required_skills, recruiter_id, created_at, updated_at, status
              FROM jobs WHERE status = $1`

	rows, err := r.db.Query(query, status)
	if err != nil {
		return nil, fmt.Errorf("repository: failed to fetch jobs by status: %w", err)
	}
	defer rows.Close()

	var jobs []*models.Job
	for rows.Next() {
		job := &models.Job{}
		err := rows.Scan(
			&job.JobID, &job.Title, &job.Description, &job.Location, &job.SalaryRange, &job.RequiredSkills,
			&job.RecruiterID, &job.CreatedAt, &job.UpdatedAt, &job.Status,
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

func (r *SQLJobRepository) UpdateJob(jobID int64, job *models.Job) error {
	query := `UPDATE jobs SET 
        title = $1, 
        description = $2, 
        location = $3, 
        salary_range = $4, 
        required_skills = $5, 
        recruiter_id = $6, 
        created_at = $7, 
        updated_at = $8, 
        status = $9
        WHERE job_id = $10`
	result, err := r.db.Exec(
		query,
		job.Title, job.Description, job.Location, job.SalaryRange, job.RequiredSkills,
		job.RecruiterID, job.CreatedAt, job.UpdatedAt, job.Status, jobID,
	)
	if err != nil {
		return fmt.Errorf("repository: failed to update job: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository: failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *SQLJobRepository) DesactivateJob(jobID int64) error {
	query := `UPDATE jobs SET 
        status = $1, 
        updated_at = $2 
        WHERE job_id = $3`
	result, err := r.db.Exec(
		query,
		"closed", time.Now(), jobID,
	)
	if err != nil {
		return fmt.Errorf("repository: failed to update job status to closed: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository: failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *SQLJobRepository) RepostJob(jobID int64) error {
	query := `UPDATE jobs SET 
        status = $1, 
        updated_at = $2 
        WHERE job_id = $3`
	result, err := r.db.Exec(
		query,
		"open", time.Now(), jobID,
	)
	if err != nil {
		return fmt.Errorf("repository: failed to update job status to open: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository: failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *SQLJobRepository) DeleteJob(jobID int64) error {
	query := "DELETE FROM jobs WHERE user_id = $1"
	result, err := r.db.Exec(query, jobID)
	if err != nil {
		return fmt.Errorf("repository: failed to delete job: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository: failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
