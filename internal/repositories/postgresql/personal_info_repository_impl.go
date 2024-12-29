package postgresql

import (
	"database/sql"
	"dz-jobs-api/internal/models"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces"
	"fmt"
	"github.com/google/uuid"
)

type SQLCandidatePersonalInfoRepository struct {
	db *sql.DB
}

func NewCandidatePersonalInfoRepository(db *sql.DB) repositoryInterfaces.CandidatePersonalInfoRepository {
	return &SQLCandidatePersonalInfoRepository{
		db: db,
	}
}

func (r *SQLCandidatePersonalInfoRepository) CreatePersonalInfo(info *models.CandidatePersonalInfo) error {
	query := `
		INSERT INTO candidate_personal_info (candidate_id, name, email, phone, address, date_of_birth, gender, bio)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.Exec(query, info.CandidateID, info.Name, info.Email, info.Phone, info.Address, info.DateOfBirth, info.Gender, info.Bio)
	if err != nil {
		return fmt.Errorf("unable to create personal info: %w", err)
	}
	return nil
}

func (r *SQLCandidatePersonalInfoRepository) GetPersonalInfo(id uuid.UUID) (*models.CandidatePersonalInfo, error) {
	var info models.CandidatePersonalInfo
	query := `
		SELECT candidate_id, name, email, phone, address, date_of_birth, gender, bio
		FROM candidate_personal_info
		WHERE candidate_id = $1`
	err := r.db.QueryRow(query, id).Scan(&info.CandidateID, &info.Name, &info.Email, &info.Phone, &info.Address, &info.DateOfBirth, &info.Gender, &info.Bio)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.CandidatePersonalInfo{}, fmt.Errorf("personal info not found: %w", err)
		}
		return &models.CandidatePersonalInfo{}, fmt.Errorf("unable to fetch personal info: %w", err)
	}
	return &info, nil
}

func (r *SQLCandidatePersonalInfoRepository) UpdatePersonalInfo(info *models.CandidatePersonalInfo) error {
	query := `
		UPDATE candidate_personal_info
		SET name = $1, email = $2, phone = $3, address = $4, date_of_birth = $5, gender = $6, bio = $7
		WHERE candidate_id = $8`
	_, err := r.db.Exec(query, info.Name, info.Email, info.Phone, info.Address, info.DateOfBirth, info.Gender, info.Bio, info.CandidateID)
	if err != nil {
		return fmt.Errorf("unable to update personal info: %w", err)
	}
	return nil
}

func (r *SQLCandidatePersonalInfoRepository) DeletePersonalInfo(id uuid.UUID) error {
	query := `DELETE FROM candidate_personal_info WHERE candidate_id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("unable to delete personal info: %w", err)
	}
	return nil
}
