package interfaces

import (
	"context"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type CandidateCertificationsRepository interface {
	CreateCertification(ctx context.Context, certification *models.CandidateCertification) error
	GetCertifications(ctx context.Context, certificationID uuid.UUID) ([]models.CandidateCertification, error)
	DeleteCertification(ctx context.Context, certificationID uuid.UUID, certificationName string) error
	ValidateCertificationOwnership(ctx context.Context, certificationID uuid.UUID, certificationName string) error
}
