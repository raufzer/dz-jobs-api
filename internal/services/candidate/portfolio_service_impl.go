package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	models "dz-jobs-api/internal/models/candidate"
	interfaces "dz-jobs-api/internal/repositories/interfaces/candidate"
	"dz-jobs-api/pkg/utils"
	"net/http"

	"github.com/google/uuid"
)

type candidatePortfolioService struct {
	portfolioRepo interfaces.CandidatePortfolioRepository
}

func NewCandidatePortfolioService(repo interfaces.CandidatePortfolioRepository) *candidatePortfolioService {
	return &candidatePortfolioService{portfolioRepo: repo}
}

func (s *candidatePortfolioService) AddProject(candidateID uuid.UUID, request request.AddProjectRequest) (*models.CandidatePortfolio, error) {
	portfolio := &models.CandidatePortfolio{
		ProjectID:   uuid.New(),
		CandidateID: candidateID,
		ProjectName: request.ProjectName,
		ProjectLink: request.ProjectLink,
		Category:    request.Category,
		Description: request.Description,
	}

	err := s.portfolioRepo.CreateProject(portfolio)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to add portfolio project")
	}

	return portfolio, nil
}

func (s *candidatePortfolioService) GetPortfolioByCandidateID(candidateID uuid.UUID) ([]models.CandidatePortfolio, error) {
	portfolio, err := s.portfolioRepo.GetPortfolioByCandidateID(candidateID)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusNotFound, "No portfolio projects found")
	}

	return portfolio, nil
}

func (s *candidatePortfolioService) DeleteProject(projectID uuid.UUID, projectName string) error {
	err := s.portfolioRepo.DeleteProject(projectID, projectName)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete portfolio project")
	}

	return nil
}
