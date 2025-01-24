package postgresql

import (
	"database/sql"

	"dz-jobs-api/internal/models"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces"

	"fmt"

	"github.com/google/uuid"
)

type SQLBookmarksRepository struct {
	db *sql.DB
}

func NewBookmarskRepository(db *sql.DB) repositoryInterfaces.BookmarksRepository {
	return &SQLBookmarksRepository{
		db: db,
	}
}

func (r *SQLBookmarksRepository) AddBookmark(candidateID uuid.UUID, jobID int64) error {
	query := "INSERT INTO bookmarks (candidate_id, job_id) VALUES ($1, $2)"
	_, err := r.db.Exec(query, candidateID, jobID)
	if err != nil {
		return fmt.Errorf("repository: failed to add bookmark: %w", err)
	}
	return nil
}

func (r *SQLBookmarksRepository) RemoveBookmark(candidateID uuid.UUID, jobID int64) error {
	query := "DELETE FROM bookmarks WHERE candidate_id = $1 AND job_id = $2"
	_, err := r.db.Exec(query, candidateID, jobID)
	if err != nil {
		return fmt.Errorf("repository: failed to remove bookmark: %w", err)
	}
	return nil
}

func (r *SQLBookmarksRepository) GetBookmarks(candidateID uuid.UUID) ([]*models.Job, error) {
	query := `
        SELECT j.job_id, j.title, j.description, j.location, j.salary_range, j.required_skills, j.recruiter_id, j.created_at, j.updated_at, j.status
        FROM bookmarks b
        JOIN jobs j ON b.job_id = j.job_id
        WHERE b.candidate_id = $1
    `
	rows, err := r.db.Query(query, candidateID)
	if err != nil {
		return nil, fmt.Errorf("repository: failed to fetch bookmarks: %w", err)
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
		jobs = append(jobs, job)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("repository: rows error: %w", err)
	}

	return jobs, nil
}
