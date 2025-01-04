package interfaces

import (
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type CandidateEducationService interface {
	AddEducation(candidateID uuid.UUID, request request.AddEducationRequest) (*models.CandidateEducation, error)
	GetEducation(candidateID uuid.UUID) ([]models.CandidateEducation, error)
	DeleteEducation(candidateID uuid.UUID, educationID uuid.UUID) error
}
