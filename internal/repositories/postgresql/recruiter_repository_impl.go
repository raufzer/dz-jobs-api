package postgresql

import (
	"database/sql"
	"dz-jobs-api/internal/models"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type SQLRecruiterRepository struct {
	db *sql.DB
}

func NewRecruiterRepository(db *sql.DB) repositoryInterfaces.RecruiterRepository {
	return &SQLRecruiterRepository{
		db: db,
	}
}

func (r *SQLRecruiterRepository) CreateRecruiter(recruiter *models.Recruiter) error {
	query := `INSERT INTO recruiters (recruiter_id, company_name, company_logo, company_description, 
			 company_website, company_location, company_contact, social_links, verified_status)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := r.db.Exec(query, recruiter.ID, recruiter.CompanyName, recruiter.CompanyLogo, recruiter.CompanyDescription,
		recruiter.CompanyWebsite, recruiter.CompanyLocation, recruiter.CompanyContact, recruiter.SocialLinks, recruiter.VerifiedStatus)
	if err != nil {
		return fmt.Errorf("repository: failed to create recruiter: %w", err)
	}

	return nil
}

func (r *SQLRecruiterRepository) GetRecruiter(recruiterID uuid.UUID) (*models.Recruiter, error) {

	query := `SELECT recruiter_id, company_name, company_logo, company_description, company_website, 
			  company_location, company_contact, social_links, verified_status
			  FROM recruiters WHERE recruiter_id = $1`

	row := r.db.QueryRow(query, recruiterID)
	recruiter := &models.Recruiter{}
	err := row.Scan(&recruiter.ID, &recruiter.CompanyName, &recruiter.CompanyLogo,
		&recruiter.CompanyDescription, &recruiter.CompanyWebsite, &recruiter.CompanyLocation,
		&recruiter.CompanyContact, &recruiter.SocialLinks, &recruiter.VerifiedStatus)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("repository: failed to fetch recruiter by ID: %w", err)
	}
	return recruiter, nil
}

func (r *SQLRecruiterRepository) UpdateRecruiter(recruiterID uuid.UUID, recruiter *models.Recruiter) error {

	query := `UPDATE recruiters SET company_name = $1, company_logo = $2, company_description = $3, 
			  company_website = $4, company_location = $5, company_contact = $6, social_links = $7, 
			  verified_status = $8 WHERE recruiter_id = $9`
	result, err := r.db.Exec(query, recruiter.CompanyName, recruiter.CompanyLogo, recruiter.CompanyDescription,
		recruiter.CompanyWebsite, recruiter.CompanyLocation, recruiter.CompanyContact, recruiter.SocialLinks, recruiter.VerifiedStatus, recruiterID)
	if err != nil {
		return fmt.Errorf("repository: failed to update recruiter: %w", err)
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

func (r *SQLRecruiterRepository) DeleteRecruiter(recruiterID uuid.UUID) error {

	query := `DELETE FROM recruiters WHERE recruiter_id = $1`
	result, err := r.db.Exec(query, recruiterID)
	if err != nil {
		return fmt.Errorf("repository: failed to delete recruiter: %w", err)
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
