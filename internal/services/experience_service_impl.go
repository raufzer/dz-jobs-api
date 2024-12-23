package services

import (
    "dz-jobs-api/internal/dto/request"
	"dz-jobs-api/pkg/utils"
    "dz-jobs-api/internal/models"
    "dz-jobs-api/internal/repositories/interfaces"
	"net/http"

	"github.com/google/uuid"
)

type CandidateExperienceService struct {
	candidateExperienceRepo interfaces.CandidateExperienceRepository
}

func NewCandidateExperienceService(repo interfaces.CandidateExperienceRepository) *CandidateExperienceService {
	return &CandidateExperienceService{candidateExperienceRepo: repo}
}

func (s *CandidateExperienceService) AddExperience(candidateID uuid.UUID, request request.AddExperienceRequest) (*models.CandidateExperience, error) {
	experience := &models.CandidateExperience{
		ExperienceID: uuid.New(),
		CandidateID:  candidateID,
		JobTitle:     request.JobTitle,
		Company:      request.Company,
		StartDate:    request.StartDate,
		EndDate:      request.EndDate,
		Description:  request.Description,
	}

	err := s.candidateExperienceRepo.CreateExperience(experience)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to add experience")
	}

	return experience, nil
}

func (s *CandidateExperienceService) GetExperience(candidateID uuid.UUID) ([]models.CandidateExperience, error) {
	experiences, err := s.candidateExperienceRepo.GetExperience(candidateID)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusNotFound, "No experience records found")
	}

	return experiences, nil
}

func (s *CandidateExperienceService) DeleteExperience(experienceID uuid.UUID) error {
	err := s.candidateExperienceRepo.DeleteExperience(experienceID)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete experience")
	}

	return nil
}
