package services

import (
	"context"
	"database/sql"
	"dz-jobs-api/config"
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"
	"dz-jobs-api/internal/repositories/interfaces"
	"dz-jobs-api/pkg/utils"
	"net/http"

	"github.com/google/uuid"
)

type CandidateEducationService struct {
	candidateEducationRepo interfaces.CandidateEducationRepository
	config                 *config.AppConfig
}

func NewCandidateEducationService(repo interfaces.CandidateEducationRepository, config *config.AppConfig) *CandidateEducationService {
	return &CandidateEducationService{
		candidateEducationRepo: repo,
		config:                 config,
	}
}

func (s *CandidateEducationService) AddEducation(ctx context.Context, candidateID uuid.UUID, request request.AddEducationRequest) (*models.CandidateEducation, error) {

	education := &models.CandidateEducation{
		ID:          uuid.New(),
		CandidateID: candidateID,
		Degree:      request.Degree,
		Institution: request.Institution,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		Description: request.Description,
	}

	err := s.candidateEducationRepo.CreateEducation(ctx, education)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to add education")
	}

	return education, nil
}

func (s *CandidateEducationService) GetEducation(ctx context.Context, candidateID uuid.UUID) ([]models.CandidateEducation, error) {
	educations, err := s.candidateEducationRepo.GetEducation(ctx, candidateID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "No education records found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch education records")
	}

	return educations, nil
}

func (s *CandidateEducationService) DeleteEducation(ctx context.Context, candidateID uuid.UUID, educationID uuid.UUID) error {
	err := s.candidateEducationRepo.DeleteEducation(ctx, candidateID, educationID)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete education")
	}

	return nil
}
