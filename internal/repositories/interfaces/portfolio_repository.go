package interfaces

import (
	"context"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type CandidatePortfolioRepository interface {
	CreateProject(ctx context.Context, project *models.CandidatePortfolio) error
	GetPortfolio(ctx context.Context, candidateID uuid.UUID) ([]models.CandidatePortfolio, error)
	DeleteProject(ctx context.Context, projectID uuid.UUID, projectName string) error
}
