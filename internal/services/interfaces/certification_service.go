package interfaces

import (
    "context"
    "dz-jobs-api/internal/dto/request"
    "dz-jobs-api/internal/models"

    "github.com/google/uuid"
)

type CandidateCertificationsService interface {
    AddCertification(ctx context.Context, candidateID uuid.UUID, request request.AddCertificationRequest) (*models.CandidateCertification, error)
    GetCertifications(ctx context.Context, candidateID uuid.UUID) ([]models.CandidateCertification, error)
    DeleteCertification(ctx context.Context, certificationID uuid.UUID, certificationName string) error
}