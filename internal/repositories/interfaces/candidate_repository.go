package interfaces

import (
	"dz-jobs-api/internal/models"
	"github.com/google/uuid"
)

type CandidateRepository interface {
	CreateCandidate(candidate *models.Candidate) (uuid.UUID, error)
	GetCandidate(candidateID uuid.UUID) (*models.Candidate, error)
	UpdateCandidate(candidateID uuid.UUID, candidate *models.Candidate) error
	DeleteCandidate(candidateID uuid.UUID) error
}
