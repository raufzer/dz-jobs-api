package services

import (
    "context" // Add this import
    "database/sql"
    "dz-jobs-api/internal/dto/request"
    "dz-jobs-api/internal/models"
    "dz-jobs-api/internal/repositories/interfaces"
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

func (s *CandidatePortfolioService) AddProject(ctx context.Context, candidateID uuid.UUID, request request.AddProjectRequest) (*models.CandidatePortfolio, error) {
    portfolio := &models.CandidatePortfolio{
        ID:          uuid.New(),
        CandidateID: candidateID,
        ProjectName: request.ProjectName,
        ProjectLink: request.ProjectLink,
        Category:    request.Category,
        Description: request.Description,
    }

    err := s.candidatePortfolioRepo.CreateProject(ctx, portfolio) // Pass context
    if err != nil {
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to add portfolio project")
    }

    return portfolio, nil
}

func (s *CandidatePortfolioService) GetPortfolio(ctx context.Context, candidateID uuid.UUID) ([]models.CandidatePortfolio, error) {
    portfolio, err := s.candidatePortfolioRepo.GetPortfolio(ctx, candidateID) // Pass context
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, utils.NewCustomError(http.StatusNotFound, "No portfolio projects found")
        }
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch portfolio projects")
    }

    return portfolio, nil
}

func (s *CandidatePortfolioService) DeleteProject(ctx context.Context, projectID uuid.UUID, projectName string) error {
    err := s.candidatePortfolioRepo.DeleteProject(ctx, projectID, projectName) // Pass context
    if err != nil {
        if err == sql.ErrNoRows {
            return utils.NewCustomError(http.StatusNotFound, "Portfolio project not found")
        }
        return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete portfolio project")
    }

    return nil
}