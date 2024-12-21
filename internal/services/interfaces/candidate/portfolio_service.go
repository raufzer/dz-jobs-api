package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	models "dz-jobs-api/internal/models/candidate"

	"github.com/google/uuid"
)

type CandidatePortfolioService interface {
	AddProject(candidateID uuid.UUID, request request.AddProjectRequest) (*models.CandidatePortfolio, error)
	GetPortfolioByCandidateID(candidateID uuid.UUID) ([]models.CandidatePortfolio, error)
	DeleteProject(projectID uuid.UUID, projectName string) error
}
