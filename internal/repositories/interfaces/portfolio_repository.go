package interfaces

import (
	"dz-jobs-api/internal/models"
	"github.com/google/uuid"
)

type CandidatePortfolioRepository interface {
	CreateProject(project *models.CandidatePortfolio) error
	GetPortfolio(id uuid.UUID) ([]models.CandidatePortfolio, error)
	DeleteProject(id uuid.UUID, projectName string) error
}
