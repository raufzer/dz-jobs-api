package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	models "dz-jobs-api/internal/models/candidate"
	interfaces "dz-jobs-api/internal/repositories/interfaces/candidate"
	"dz-jobs-api/pkg/utils"
	"net/http"

	"github.com/google/uuid"
)

type CandidatePortfolioService struct {
	candidatePortfolioRepo interfaces.CandidatePortfolioRepository
}

func NewCandidatePortfolioService(repo interfaces.CandidatePortfolioRepository) *CandidatePortfolioService {
	return &CandidatePortfolioService{candidatePortfolioRepo: repo}
}

func (s *CandidatePortfolioService) AddProject(candidateID uuid.UUID, request request.AddProjectRequest) (*models.CandidatePortfolio, error) {
	portfolio := &models.CandidatePortfolio{
		ProjectID:   uuid.New(),
		CandidateID: candidateID,
		ProjectName: request.ProjectName,
		ProjectLink: request.ProjectLink,
		Category:    request.Category,
		Description: request.Description,
	}

	err := s.candidatePortfolioRepo.CreateProject(portfolio)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to add portfolio project")
	}

	return portfolio, nil
}

func (s *CandidatePortfolioService) GetPortfolio(candidateID uuid.UUID) ([]models.CandidatePortfolio, error) {
	portfolio, err := s.candidatePortfolioRepo.GetPortfolio(candidateID)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusNotFound, "No portfolio projects found")
	}

	return portfolio, nil
}

func (s *CandidatePortfolioService) DeleteProject(projectID uuid.UUID, projectName string) error {
	err := s.candidatePortfolioRepo.DeleteProject(projectID, projectName)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete portfolio project")
	}

	return nil
}
