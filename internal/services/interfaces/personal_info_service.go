package interfaces

import (
    "context"
    "dz-jobs-api/internal/dto/request"
    "dz-jobs-api/internal/models"

    "github.com/google/uuid"
)

type CandidatePersonalInfoService interface {
    AddPersonalInfo(ctx context.Context, request request.AddPersonalInfoRequest, candidateID uuid.UUID) (*models.CandidatePersonalInfo, error)
    UpdatePersonalInfo(ctx context.Context, id uuid.UUID, request request.UpdatePersonalInfoRequest) (*models.CandidatePersonalInfo, error)
    GetPersonalInfo(ctx context.Context, candidateID uuid.UUID) (*models.CandidatePersonalInfo, error)
    DeletePersonalInfo(ctx context.Context, candidateID uuid.UUID) error
}