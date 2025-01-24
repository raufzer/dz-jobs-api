package interfaces

import (
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type RecruiterRepository interface {
	CreateRecruiter(recruiter *models.Recruiter) error
	GetRecruiter(recruiterID uuid.UUID) (*models.Recruiter, error)
	UpdateRecruiter(recruiterID uuid.UUID, recruiter *models.Recruiter) error
	DeleteRecruiter(recruiterID uuid.UUID) error
}
