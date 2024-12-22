package candidate

import (
	"database/sql"
	models "dz-jobs-api/internal/models/candidate"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces/candidate"
	"fmt"
	"github.com/google/uuid"
)

type SQLCandidateExperienceRepository struct {
	db *sql.DB
}

func NewCandidateExperienceRepository(db *sql.DB) repositoryInterfaces.CandidateExperienceRepository {
	return &SQLCandidateExperienceRepository{
		db: db,
	}
}

func (r *SQLCandidateExperienceRepository) CreateExperience(experience *models.CandidateExperience) error {
	query := `INSERT INTO candidate_experience (experience_id, candidate_id, job_title, company, start_date, end_date, description) 
			VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, experience.ExperienceID, experience.CandidateID, experience.JobTitle, experience.Company, experience.StartDate, experience.EndDate, experience.Description)
	if err != nil {
		return fmt.Errorf("repository: failed to create experience: %w", err)
	}
	return nil
}

func (r *SQLCandidateExperienceRepository) GetExperience(id uuid.UUID) ([]models.CandidateExperience, error) {
	rows, err := r.db.Query(`SELECT experience_id, candidate_id, job_title, company, start_date, end_date, description FROM candidate_experience WHERE candidate_id = $1`, id)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch experience: %w", err)
	}
	defer rows.Close()

	var experiences []models.CandidateExperience
	for rows.Next() {
		var experience models.CandidateExperience
		if err := rows.Scan(&experience.ExperienceID, &experience.CandidateID, &experience.JobTitle, &experience.Company, &experience.StartDate, &experience.EndDate, &experience.Description); err != nil {
			return nil, fmt.Errorf("unable to scan experience data: %w", err)
		}
		experiences = append(experiences, experience)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}
	return experiences, nil
}

func (r *SQLCandidateExperienceRepository) DeleteExperience(id uuid.UUID) error {
	query := `DELETE FROM candidate_experience WHERE experience_id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("unable to delete experience: %w", err)
	}
	return nil
}
