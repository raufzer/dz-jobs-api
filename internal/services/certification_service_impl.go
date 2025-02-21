package services

import (
	"context"
	"database/sql"
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"
	"dz-jobs-api/internal/repositories/interfaces"
	"dz-jobs-api/pkg/utils"
	"net/http"

	"github.com/google/uuid"
)

type CandidateCertificationsService struct {
    candidateCertificationsRepo interfaces.CandidateCertificationsRepository
}

func NewCandidateCertificationsService(repo interfaces.CandidateCertificationsRepository) *CandidateCertificationsService {
    return &CandidateCertificationsService{candidateCertificationsRepo: repo}
}

func (s *CandidateCertificationsService) AddCertification(ctx context.Context, candidateID uuid.UUID, request request.AddCertificationRequest) (*models.CandidateCertification, error) {
    certification := &models.CandidateCertification{
        ID:                uuid.New(),
        CandidateID:       candidateID,
        CertificationName: request.CertificationName,
        IssuedBy:          request.IssuedBy,
        IssueDate:         request.IssueDate,
        ExpirationDate:    request.ExpirationDate,
    }

    err := s.candidateCertificationsRepo.CreateCertification(ctx, certification)
    if err != nil {
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to add certification")
    }

    return certification, nil
}

func (s *CandidateCertificationsService) GetCertifications(ctx context.Context, candidateID uuid.UUID) ([]models.CandidateCertification, error) {
    certifications, err := s.candidateCertificationsRepo.GetCertifications(ctx, candidateID)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, utils.NewCustomError(http.StatusNotFound, "No certifications found")
        }
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch certifications")
    }

    return certifications, nil
}

func (s *CandidateCertificationsService) DeleteCertification(ctx context.Context, certificationID uuid.UUID, certificationName string) error {
    err := s.candidateCertificationsRepo.ValidateCertificationOwnership(ctx, certificationID, certificationName)
    if err != nil {
        return utils.NewCustomError(http.StatusForbidden, "You do not own this certification")
    }
    err = s.candidateCertificationsRepo.DeleteCertification(ctx, certificationID, certificationName)
    if err != nil {
        return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete certification")
    }

    return nil
}