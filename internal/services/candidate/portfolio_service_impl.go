package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	"dz-jobs-api/internal/helpers"
	models "dz-jobs-api/internal/models/candidate"
	interfaces "dz-jobs-api/internal/repositories/interfaces/candidate"
	"net/http"

	"github.com/google/uuid"
)

type candidatePortfolioService struct {
	portfolioRepo interfaces.CandidatePortfolioRepository
}

func NewCandidatePortfolioService(repo interfaces.CandidatePortfolioRepository) *candidatePortfolioService {
	return &candidatePortfolioService{portfolioRepo: repo}
}

func (s *candidatePortfolioService) AddPortfolio(request request.AddPortfolioRequest) (*models.CandidatePortfolio, error) {
	portfolio := &models.CandidatePortfolio{
		ProjectID:   uuid.New(),
		ProjectName: request.ProjectName,
		ProjectLink: request.ProjectLink,
		Category:    request.Category,
		Description: request.Description,
	}

	err := s.portfolioRepo.CreatePortfolio(*portfolio)
	if err != nil {
		return nil, helpers.NewCustomError(http.StatusInternalServerError, "Failed to add portfolio project")
	}

	return portfolio, nil
}

func (s *candidatePortfolioService) GetPortfolioByCandidateID(candidateID uuid.UUID) ([]models.CandidatePortfolio, error) {
	portfolio, err := s.portfolioRepo.GetPortfolioByCandidateID(candidateID)
	if err != nil {
		return nil, helpers.NewCustomError(http.StatusNotFound, "No portfolio projects found")
	}

	return portfolio, nil
}

func (s *candidatePortfolioService) DeletePortfolio(projectID uuid.UUID, projectName string) error {
	err := s.portfolioRepo.DeletePortfolio(projectID, projectName)
	if err != nil {
		return helpers.NewCustomError(http.StatusInternalServerError, "Failed to delete portfolio project")
	}

	return nil
}
