package postgresql

import (
		"context"
	"database/sql"
	"dz-jobs-api/internal/models"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces"
	"fmt"

	"github.com/google/uuid"
)

type SQLCandidateEducationRepository struct {
	db *sql.DB
}

func NewCandidateEducationRepository(db *sql.DB) repositoryInterfaces.CandidateEducationRepository {
	return &SQLCandidateEducationRepository{
		db: db,
	}
}

func (r *SQLCandidateEducationRepository) CreateEducation(ctx context.Context, education *models.CandidateEducation) error {
	query := "INSERT INTO candidate_education (education_id, candidate_id, degree, institution, start_date, end_date, description) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := r.db.Exec(query, education.ID, education.CandidateID, education.Degree, education.Institution, education.StartDate, education.EndDate, education.Description)
	if err != nil {
		return fmt.Errorf("repository: failed to create education: %w", err)
	}
	return nil
}

func (r *SQLCandidateEducationRepository) GetEducation(ctx context.Context, educationID uuid.UUID) ([]models.CandidateEducation, error) {
	rows, err := r.db.Query(`SELECT education_id, candidate_id, degree, institution, start_date, end_date, description FROM candidate_education WHERE candidate_id = $1`, educationID)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch education: %w", err)
	}
	defer rows.Close()

	var educations []models.CandidateEducation
	for rows.Next() {
		var education models.CandidateEducation
		if err := rows.Scan(&education.ID, &education.CandidateID, &education.Degree, &education.Institution, &education.StartDate, &education.EndDate, &education.Description); err != nil {
			return nil, fmt.Errorf("unable to scan education data: %w", err)
		}
		educations = append(educations, education)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}
	return educations, nil
}

func (r *SQLCandidateEducationRepository) DeleteEducation(ctx context.Context, candidateID, educationID uuid.UUID) error {
	query := `DELETE FROM candidate_education WHERE education_id = $1 AND candidate_id = $2`
	_, err := r.db.Exec(query, educationID, candidateID)
	if err != nil {
		return fmt.Errorf("unable to delete education: %w", err)
	}
	return nil
}
