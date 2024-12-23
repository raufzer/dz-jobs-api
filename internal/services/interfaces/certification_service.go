package interfaces

import (
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type CandidateCertificationsService interface {
	AddCertification(candidateID uuid.UUID, request request.AddCertificationRequest) (*models.CandidateCertification, error)
	GetCertifications(candidateID uuid.UUID) ([]models.CandidateCertification, error)
	DeleteCertification(certificationID uuid.UUID, certificationName string) error
}
