package interfaces

import (
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type CandidatePortfolioService interface {
	AddProject(candidateID uuid.UUID, request request.AddProjectRequest) (*models.CandidatePortfolio, error)
	GetPortfolio(candidateID uuid.UUID) ([]models.CandidatePortfolio, error)
	DeleteProject(projectID uuid.UUID, projectName string) error
}
