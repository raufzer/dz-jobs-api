package postgresql

import (
	"database/sql"
	"dz-jobs-api/internal/models"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces"
	"fmt"
	"github.com/google/uuid"
)

type SQLCandidateRepository struct {
	db *sql.DB
}

func NewCandidateRepository(db *sql.DB) repositoryInterfaces.CandidateRepository {
	return &SQLCandidateRepository{
		db: db,
	}
}

func (r *SQLCandidateRepository) CreateCandidate(candidate *models.Candidate) (uuid.UUID, error) {
	query := `INSERT INTO candidates (candidate_id, resume, profile_picture) VALUES ($1, $2, $3) RETURNING candidate_id`

	err := r.db.QueryRow(query, candidate.CandidateID, candidate.Resume, candidate.ProfilePicture).Scan(&candidate.CandidateID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("unable to create candidate: %w", err)
	}

	return candidate.CandidateID, nil
}

func (r *SQLCandidateRepository) GetCandidate(id uuid.UUID) (*models.Candidate, error) {
	query := `SELECT candidate_id, resume, profile_picture FROM candidates WHERE candidate_id = $1`
	row := r.db.QueryRow(query, id)
	candidate := &models.Candidate{}
	err := row.Scan(&candidate.CandidateID, &candidate.Resume, &candidate.ProfilePicture)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("candidate not found: %w", err)
		}
		return nil, fmt.Errorf("unable to fetch candidate: %w", err)
	}
	return candidate, nil
}

func (r *SQLCandidateRepository) UpdateCandidate(candidate_id uuid.UUID, candidate *models.Candidate) error {
	query := `UPDATE candidates SET resume = $1, profile_picture = $2 WHERE candidate_id = $3`
	result, err := r.db.Exec(query, candidate.Resume, candidate.ProfilePicture, candidate_id)
	if err != nil {
		return fmt.Errorf("repository: failed to update user: %w", err)
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

func (r *SQLCandidateRepository) DeleteCandidate(id uuid.UUID) error {
	query := `DELETE FROM candidates WHERE candidate_id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("unable to delete candidate: %w", err)
	}
	return nil
}
