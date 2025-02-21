package interfaces

import (
    "context"
    "dz-jobs-api/internal/dto/request"
    "dz-jobs-api/internal/models"

    "github.com/google/uuid"
)

type CandidateEducationService interface {
    AddEducation(ctx context.Context, candidateID uuid.UUID, request request.AddEducationRequest) (*models.CandidateEducation, error)
    GetEducation(ctx context.Context, candidateID uuid.UUID) ([]models.CandidateEducation, error)
    DeleteEducation(ctx context.Context, candidateID uuid.UUID, educationID uuid.UUID) error
}