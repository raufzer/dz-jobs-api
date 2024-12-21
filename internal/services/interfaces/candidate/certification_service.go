package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	models "dz-jobs-api/internal/models/candidate"
	"github.com/google/uuid"
)

type CandidateCertificationsService interface {
	AddCertification(candidateID uuid.UUID, request request.AddCertificationRequest) (*models.CandidateCertification, error)
	GetCertificationsByCandidateID(candidateID uuid.UUID) ([]models.CandidateCertification, error)
	DeleteCertification(certificationID uuid.UUID, certificationName string) error
}
