package interfaces

import (
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type RecruiterRepository interface {
	CreateRecruiter(recruiter *models.Recruiter) error
	GetRecruiterByID(recruiter_id uuid.UUID) (*models.Recruiter, error)
	UpdateRecruiter(recruiter_id uuid.UUID, recruiter *models.Recruiter) error
	DeleteRecruiter(recruiter_id uuid.UUID) error
}
