package interfaces

import (
	"dz-jobs-api/internal/models"
	"github.com/google/uuid"
)

type CandidateEducationRepository interface {
	CreateEducation(education *models.CandidateEducation) error
	GetEducation(id uuid.UUID) ([]models.CandidateEducation, error)
	DeleteEducation(id uuid.UUID) error
}
