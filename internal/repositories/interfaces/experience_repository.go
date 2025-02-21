package interfaces

import (
	"context"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type CandidateExperienceRepository interface {
	CreateExperience(ctx context.Context, experience *models.CandidateExperience) error
	GetExperience(ctx context.Context, experienceID uuid.UUID) ([]models.CandidateExperience, error)
	DeleteExperience(ctx context.Context, candidateID uuid.UUID, experienceID uuid.UUID) error
}
