package interfaces

import (
    "context"
    "dz-jobs-api/internal/dto/request"
    "dz-jobs-api/internal/models"

    "github.com/google/uuid"
)

type CandidateExperienceService interface {
    AddExperience(ctx context.Context, candidateID uuid.UUID, request request.AddExperienceRequest) (*models.CandidateExperience, error)
    GetExperience(ctx context.Context, candidateID uuid.UUID) ([]models.CandidateExperience, error)
    DeleteExperience(ctx context.Context, candidateID uuid.UUID, experienceID uuid.UUID) error
}