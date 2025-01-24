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

	err := r.db.QueryRow(query, candidate.ID, candidate.Resume, candidate.ProfilePicture).Scan(&candidate.ID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("unable to create candidate: %w", err)
	}

	return candidate.ID, nil
}

func (r *SQLCandidateRepository) GetCandidate(candidateID uuid.UUID) (*models.Candidate, error) {
	query := `SELECT candidate_id, resume, profile_picture FROM candidates WHERE candidate_id = $1`
	row := r.db.QueryRow(query, candidateID)
	candidate := &models.Candidate{}
	err := row.Scan(&candidate.ID, &candidate.Resume, &candidate.ProfilePicture)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("candidate not found: %w", err)
		}
		return nil, fmt.Errorf("unable to fetch candidate: %w", err)
	}
	return candidate, nil
}

func (r *SQLCandidateRepository) UpdateCandidate(candidateID uuid.UUID, candidate *models.Candidate) error {
	query := `UPDATE candidates SET resume = $1, profile_picture = $2 WHERE candidate_id = $3`
	result, err := r.db.Exec(query, candidate.Resume, candidate.ProfilePicture, candidateID)
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

func (r *SQLCandidateRepository) DeleteCandidate(candidateID uuid.UUID) error {
	query := `DELETE FROM candidates WHERE candidate_id = $1`
	_, err := r.db.Exec(query, candidateID)
	if err != nil {
		return fmt.Errorf("unable to delete candidate: %w", err)
	}
	return nil
}
