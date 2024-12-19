package candidate

import (
	models"dz-jobs-api/internal/models/candidate"
	"github.com/google/uuid"
)

type CandidatePortfolioRepository interface {
	CreatePortfolio(project models.CandidatePortfolio) error
	GetPortfolioByCandidateID(id uuid.UUID) ([]models.CandidatePortfolio, error)
	UpdatePortfolio(project models.CandidatePortfolio) error
	DeletePortfolio(id uuid.UUID, projectName string) error
}
