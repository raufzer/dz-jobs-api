package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	"dz-jobs-api/pkg/utils"
	models "dz-jobs-api/internal/models/candidate"
	interfaces "dz-jobs-api/internal/repositories/interfaces/candidate"
	"net/http"

	"github.com/google/uuid"
)

type candidateCertificationsService struct {
	certificationsRepo interfaces.CandidateCertificationsRepository
}

func NewCandidateCertificationsService(repo interfaces.CandidateCertificationsRepository) *candidateCertificationsService {
	return &candidateCertificationsService{certificationsRepo: repo}
}

func (s *candidateCertificationsService) AddCertification(candidateID uuid.UUID, request request.AddCertificationRequest) (*models.CandidateCertification, error) {
	certification := &models.CandidateCertification{
		CertificationID:   uuid.New(),
		CandidateID:       candidateID,
		CertificationName: request.CertificationName,
		IssuedBy:          request.IssuedBy,
		IssueDate:         request.IssueDate,
		ExpirationDate:    request.ExpirationDate,
	}

	err := s.certificationsRepo.CreateCertification(certification)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to add certification")
	}

	return certification, nil
}

func (s *candidateCertificationsService) GetCertificationsByCandidateID(candidateID uuid.UUID) ([]models.CandidateCertification, error) {
	certifications, err := s.certificationsRepo.GetCertificationsByCandidateID(candidateID)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusNotFound, "No certifications found")
	}

	return certifications, nil
}

func (s *candidateCertificationsService) DeleteCertification(certificationID uuid.UUID, certificationName string) error {
	err := s.certificationsRepo.DeleteCertification(certificationID, certificationName)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete certification")
	}

	return nil
}
