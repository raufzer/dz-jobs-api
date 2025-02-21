package interfaces

import (
	"context"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type CandidateEducationRepository interface {
	CreateEducation(ctx context.Context, education *models.CandidateEducation) error
	GetEducation(ctx context.Context, educationID uuid.UUID) ([]models.CandidateEducation, error)
	DeleteEducation(ctx context.Context, candidateID uuid.UUID, educationID uuid.UUID) error
}
