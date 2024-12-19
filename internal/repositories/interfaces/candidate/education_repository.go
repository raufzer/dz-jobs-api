package candidate

import (
	models"dz-jobs-api/internal/models/candidate"
	"github.com/google/uuid"
)

type CandidateEducationRepository interface {
	CreateEducation(education models.CandidateEducation) error
	GetEducationByCandidateID(id uuid.UUID) ([]models.CandidateEducation, error)
	UpdateEducation(education models.CandidateEducation) error
	DeleteEducation(id uuid.UUID) error
}
