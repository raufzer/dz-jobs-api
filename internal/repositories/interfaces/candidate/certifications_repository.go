package candidate

import (
	models "dz-jobs-api/internal/models/candidate"
	"github.com/google/uuid"
)

type CandidateCertificationsRepository interface {
	CreateCertification(certification *models.CandidateCertification) error
	GetCertificationsByCandidateID(id uuid.UUID) ([]models.CandidateCertification, error)
	UpdateCertification(certification models.CandidateCertification) error
	DeleteCertification(id uuid.UUID, certificationName string) error
}
