package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	models "dz-jobs-api/internal/models/candidate"
	"github.com/google/uuid"
)

type CandidateExperienceService interface {
	AddExperience(candidateID uuid.UUID, request request.AddExperienceRequest) (*models.CandidateExperience, error)
	GetExperienceByCandidateID(candidateID uuid.UUID) ([]models.CandidateExperience, error)
	DeleteExperience(experienceID uuid.UUID) error
}
