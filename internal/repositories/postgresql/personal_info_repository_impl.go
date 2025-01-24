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
	_, err := r.db.Exec(query, info.ID, info.Name, info.Email, info.Phone, info.Address, info.DateOfBirth, info.Gender, info.Bio)
	if err != nil {
		return fmt.Errorf("unable to create personal info: %w", err)
	}
	return nil
}
func (r *SQLCandidatePersonalInfoRepository) GetPersonalInfo(candidateID uuid.UUID) (*models.CandidatePersonalInfo, error) {
	var info models.CandidatePersonalInfo
	query := `
		SELECT candidate_id, name, email, phone, address, date_of_birth, gender, bio
		FROM candidate_personal_info
		WHERE candidate_id = $1`
	err := r.db.QueryRow(query, candidateID).Scan(&info.ID, &info.Name, &info.Email, &info.Phone, &info.Address, &info.DateOfBirth, &info.Gender, &info.Bio)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.CandidatePersonalInfo{}, fmt.Errorf("personal info not found: %w", err)
		}
		return &models.CandidatePersonalInfo{}, fmt.Errorf("unable to fetch personal info: %w", err)
	}
	return &info, nil
}
func (r *SQLCandidatePersonalInfoRepository) UpdatePersonalInfo(info *models.CandidatePersonalInfo) error {
	query := `UPDATE candidate_personal_info SET`
	args := []interface{}{}
	argIndex := 1
	
	if info.Name != "" {
		query += fmt.Sprintf(" name = $%d,", argIndex)
		args = append(args, info.Name)
		argIndex++
	}
	if info.Email != "" {
		query += fmt.Sprintf(" email = $%d,", argIndex)
		args = append(args, info.Email)
		argIndex++
	}
	if info.Phone != "" {
		query += fmt.Sprintf(" phone = $%d,", argIndex)
		args = append(args, info.Phone)
		argIndex++
	}
	if info.Address != "" {
		query += fmt.Sprintf(" address = $%d,", argIndex)
		args = append(args, info.Address)
		argIndex++
	}
	if info.DateOfBirth != "" {
		query += fmt.Sprintf(" date_of_birth = $%d,", argIndex)
		args = append(args, info.DateOfBirth)
		argIndex++
	}
	if info.Gender != "" {
		query += fmt.Sprintf(" gender = $%d,", argIndex)
		args = append(args, info.Gender)
		argIndex++
	}
	if info.Bio != "" {
		query += fmt.Sprintf(" bio = $%d,", argIndex)
		args = append(args, info.Bio)
		argIndex++
	}
	
	if len(args) == 0 {
		return fmt.Errorf("no fields to update")
	}
	
	query = query[:len(query)-1]
	query += fmt.Sprintf(" WHERE candidate_id = $%d", argIndex)
	args = append(args, info.ID)
	
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("unable to update personal info: %w", err)
	}
	return nil
}
func (r *SQLCandidatePersonalInfoRepository) DeletePersonalInfo(candidateID uuid.UUID) error {
	query := `DELETE FROM candidate_personal_info WHERE candidate_id = $1`
	_, err := r.db.Exec(query, candidateID)
	if err != nil {
		return fmt.Errorf("unable to delete personal info: %w", err)
	}
	return nil
}
