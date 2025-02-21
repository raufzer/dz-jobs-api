package interfaces

import (
	"context"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type CandidateRepository interface {
	CreateCandidate(ctx context.Context, candidate *models.Candidate) (uuid.UUID, error)
	GetCandidate(ctx context.Context, candidateID uuid.UUID) (*models.Candidate, error)
	UpdateCandidate(ctx context.Context, candidateID uuid.UUID, candidate *models.Candidate) error
	DeleteCandidate(ctx context.Context, candidateID uuid.UUID) error
}
