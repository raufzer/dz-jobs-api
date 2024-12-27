package postgresql

import (
	"database/sql"
	"dz-jobs-api/internal/models"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces"
	"fmt"

	"github.com/google/uuid"
)

type SQLCandidateCertificationRepository struct {
	db *sql.DB
}

func NewCandidateCertificationsRepository(db *sql.DB) repositoryInterfaces.CandidateCertificationsRepository {
	return &SQLCandidateCertificationRepository{db: db}
}

func (r *SQLCandidateCertificationRepository) CreateCertification(certification *models.CandidateCertification) error {
	query := `INSERT INTO candidate_certifications (certification_id, candidate_id, certification_name, issued_by, issue_date, expiration_date) 
			VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, certification.CertificationID, certification.CandidateID, certification.CertificationName, certification.IssuedBy, certification.IssueDate, certification.ExpirationDate)
	if err != nil {
		return fmt.Errorf("unable to create certification: %w", err)
	}
	return nil
}

func (r *SQLCandidateCertificationRepository) GetCertifications(id uuid.UUID) ([]models.CandidateCertification, error) {
	rows, err := r.db.Query(`SELECT certification_id, candidate_id, certification_name, issued_by, issue_date, expiration_date FROM candidate_certifications WHERE candidate_id = $1`, id)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch certifications: %w", err)
	}
	defer rows.Close()

	var certifications []models.CandidateCertification
	for rows.Next() {
		var certification models.CandidateCertification
		if err := rows.Scan(&certification.CertificationID, &certification.CandidateID, &certification.CertificationName, &certification.IssuedBy, &certification.IssueDate, &certification.ExpirationDate); err != nil {
			return nil, fmt.Errorf("unable to scan certification data: %w", err)
		}
		if certification.CandidateID != id {
			return nil, fmt.Errorf("repository: unauthorized access, recruiter does not own certfications with ID %d", id)
		}
		certifications = append(certifications, certification)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}
	return certifications, nil
}
func (r *SQLCandidateCertificationRepository) DeleteCertification(id uuid.UUID, certificationName string) error {
	query := `DELETE FROM candidate_certifications WHERE candidate_id = $1 AND certification_name = $2`
	_, err := r.db.Exec(query, id, certificationName)
	if err != nil {
		return fmt.Errorf("unable to delete certification: %w", err)
	}
	return nil
}

func (r *SQLCandidateCertificationRepository) ValidateCertificationOwnership(id uuid.UUID, certificationName string) error {
	query := `SELECT candidate_id FROM candidate_certifications WHERE candidate_id = $1`
	row := r.db.QueryRow(query, id)

	var ownerID uuid.UUID
	if err := row.Scan(&ownerID); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("repository: certification not found: %w", err)
		}
		return fmt.Errorf("repository: failed to check job ownership: %w", err)
	}

	if ownerID != id {
		return fmt.Errorf("repository: unauthorized access, recruiter does not own the job")
	}

	return nil
}
