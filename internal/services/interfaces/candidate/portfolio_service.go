package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	models "dz-jobs-api/internal/models/candidate"
	"github.com/google/uuid"
)

type CandidatePortfolioService interface {
	AddPortfolio(request request.AddPortfolioRequest) (*models.CandidatePortfolio, error)
	GetPortfolioByCandidateID(candidateID uuid.UUID) ([]models.CandidatePortfolio, error)
	DeletePortfolio(projectID uuid.UUID, projectName string) error
}
