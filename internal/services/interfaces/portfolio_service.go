package interfaces

import (
    "context"
    "dz-jobs-api/internal/dto/request"
    "dz-jobs-api/internal/models"

    "github.com/google/uuid"
)

type CandidatePortfolioService interface {
    AddProject(ctx context.Context, candidateID uuid.UUID, request request.AddProjectRequest) (*models.CandidatePortfolio, error)
    GetPortfolio(ctx context.Context, candidateID uuid.UUID) ([]models.CandidatePortfolio, error)
    DeleteProject(ctx context.Context, projectID uuid.UUID, projectName string) error
}