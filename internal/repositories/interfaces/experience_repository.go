package interfaces

import (
	"dz-jobs-api/internal/models"
	"github.com/google/uuid"
)


type CandidateExperienceRepository interface {
	CreateExperience(experience *models.CandidateExperience) error
	GetExperience(experienceID uuid.UUID) ([]models.CandidateExperience, error)
	DeleteExperience(candidateID uuid.UUID, experienceID uuid.UUID) error
}
