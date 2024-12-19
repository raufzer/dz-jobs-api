package candidate

import (
	"database/sql"
	models "dz-jobs-api/internal/models/candidate"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces/candidate"
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

func (r *SQLCandidateEducationRepository) CreateEducation(education models.CandidateEducation) error {
	query := `INSERT INTO candidate_education (education_id, candidate_id, degree, institution, start_date, end_date, description) 
			VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, education.EducationID, education.CandidateID, education.Degree, education.Institution, education.StartDate, education.EndDate, education.Description)
	if err != nil {
		return fmt.Errorf("unable to create education: %w", err)
	}
	return nil
}

func (r *SQLCandidateEducationRepository) GetEducationByCandidateID(id uuid.UUID) ([]models.CandidateEducation, error) {
	rows, err := r.db.Query(`SELECT education_id, candidate_id, degree, institution, start_date, end_date, description FROM candidate_education WHERE candidate_id = $1`, id)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch education: %w", err)
	}
	defer rows.Close()

	var educations []models.CandidateEducation
	for rows.Next() {
		var education models.CandidateEducation
		if err := rows.Scan(&education.EducationID, &education.CandidateID, &education.Degree, &education.Institution, &education.StartDate, &education.EndDate, &education.Description); err != nil {
			return nil, fmt.Errorf("unable to scan education data: %w", err)
		}
		educations = append(educations, education)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}
	return educations, nil
}

func (r *SQLCandidateEducationRepository) UpdateEducation(education models.CandidateEducation) error {
	query := `UPDATE candidate_education SET degree = $1, institution = $2, start_date = $3, end_date = $4, description = $5 WHERE education_id = $6`
	_, err := r.db.Exec(query, education.Degree, education.Institution, education.StartDate, education.EndDate, education.Description, education.EducationID)
	if err != nil {
		return fmt.Errorf("unable to update education: %w", err)
	}
	return nil
}

func (r *SQLCandidateEducationRepository) DeleteEducation(id uuid.UUID) error {
	query := `DELETE FROM candidate_education WHERE education_id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("unable to delete education: %w", err)
	}
	return nil
}
