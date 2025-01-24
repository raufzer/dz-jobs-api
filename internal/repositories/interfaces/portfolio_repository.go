package interfaces

import (
	"dz-jobs-api/internal/models"
	"github.com/google/uuid"
)

type CandidatePortfolioRepository interface {
	CreateProject(project *models.CandidatePortfolio) error
	GetPortfolio(candidateID uuid.UUID) ([]models.CandidatePortfolio, error)
	DeleteProject(projectID uuid.UUID, projectName string) error
}
