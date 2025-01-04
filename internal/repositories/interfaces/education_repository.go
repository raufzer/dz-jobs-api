package interfaces

import (
	"dz-jobs-api/internal/models"
	"github.com/google/uuid"
)

type CandidateEducationRepository interface {
	CreateEducation(education *models.CandidateEducation) error
	GetEducation(educationID uuid.UUID) ([]models.CandidateEducation, error)
	DeleteEducation(candidateID uuid.UUID, educationID uuid.UUID) error
}
