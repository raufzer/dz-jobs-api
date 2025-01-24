package interfaces

import (
	"dz-jobs-api/internal/models"
	"github.com/google/uuid"
)

type CandidateCertificationsRepository interface {
	CreateCertification(certification *models.CandidateCertification) error
	GetCertifications(certificationID uuid.UUID) ([]models.CandidateCertification, error)
	DeleteCertification(certificationID uuid.UUID, certificationName string) error
	ValidateCertificationOwnership(certificationID uuid.UUID, certificationName string) error
}
