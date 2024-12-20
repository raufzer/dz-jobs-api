package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	"dz-jobs-api/internal/helpers"
	models "dz-jobs-api/internal/models/candidate"
	interfaces "dz-jobs-api/internal/repositories/interfaces/candidate"
	"net/http"

	"github.com/google/uuid"
)

type candidateExperienceService struct {
	experienceRepo interfaces.CandidateExperienceRepository
}

func NewCandidateExperienceService(repo interfaces.CandidateExperienceRepository) *candidateExperienceService {
	return &candidateExperienceService{experienceRepo: repo}
}

func (s *candidateExperienceService) AddExperience(request request.AddExperienceRequest) (*models.CandidateExperience, error) {
	experience := &models.CandidateExperience{
		ExperienceID: uuid.New(),
		JobTitle:     request.JobTitle,
		Company:      request.Company,
		StartDate:    request.StartDate,
		EndDate:      request.EndDate,
		Description:  request.Description,
	}

	err := s.experienceRepo.CreateExperience(*experience)
	if err != nil {
		return nil, helpers.NewCustomError(http.StatusInternalServerError, "Failed to add experience")
	}

	return experience, nil
}

func (s *candidateExperienceService) GetExperienceByCandidateID(candidateID uuid.UUID) ([]models.CandidateExperience, error) {
	experiences, err := s.experienceRepo.GetExperienceByCandidateID(candidateID)
	if err != nil {
		return nil, helpers.NewCustomError(http.StatusNotFound, "No experience records found")
	}

	return experiences, nil
}

func (s *candidateExperienceService) DeleteExperience(experienceID uuid.UUID) error {
	err := s.experienceRepo.DeleteExperience(experienceID)
	if err != nil {
		return helpers.NewCustomError(http.StatusInternalServerError, "Failed to delete experience")
	}

	return nil
}
