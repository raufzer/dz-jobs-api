package interfaces

import (
	"context"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type RecruiterRepository interface {
	CreateRecruiter(ctx context.Context, recruiter *models.Recruiter) error
	GetRecruiter(ctx context.Context, recruiterID uuid.UUID) (*models.Recruiter, error)
	UpdateRecruiter(ctx context.Context, recruiterID uuid.UUID, recruiter *models.Recruiter) error
	DeleteRecruiter(ctx context.Context, recruiterID uuid.UUID) error
}
