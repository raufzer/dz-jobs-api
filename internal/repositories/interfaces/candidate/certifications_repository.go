package candidate

import (
	models "dz-jobs-api/internal/models/candidate"
	"github.com/google/uuid"
)

type CandidateCertificationsRepository interface {
	CreateCertification(certification *models.CandidateCertification) error
	GetCertifications(id uuid.UUID) ([]models.CandidateCertification, error)
	DeleteCertification(id uuid.UUID, certificationName string) error
}
