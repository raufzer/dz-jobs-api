package interfaces

import (
	"dz-jobs-api/internal/models"
	"github.com/google/uuid"
)


type CandidateExperienceRepository interface {
	CreateExperience(experience *models.CandidateExperience) error
	GetExperience(id uuid.UUID) ([]models.CandidateExperience, error)
	DeleteExperience(id uuid.UUID) error
}
