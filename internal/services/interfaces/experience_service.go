package interfaces

import (
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type CandidateExperienceService interface {
	AddExperience(candidateID uuid.UUID, request request.AddExperienceRequest) (*models.CandidateExperience, error)
	GetExperience(candidateID uuid.UUID) ([]models.CandidateExperience, error)
	DeleteExperience(candidateID uuid.UUID, experienceID uuid.UUID) error
}
