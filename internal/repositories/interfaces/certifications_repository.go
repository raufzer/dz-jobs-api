package interfaces

import (
	"dz-jobs-api/internal/models"
	"github.com/google/uuid"
)

type CandidateCertificationsRepository interface {
	CreateCertification(certification *models.CandidateCertification) error
	GetCertifications(id uuid.UUID) ([]models.CandidateCertification, error)
	DeleteCertification(id uuid.UUID, certificationName string) error
}
