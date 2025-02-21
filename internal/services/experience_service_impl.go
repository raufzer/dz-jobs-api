package services

import (
    "context" // Add this import
    "database/sql"
    "dz-jobs-api/internal/dto/request"
    "dz-jobs-api/internal/models"
    "dz-jobs-api/internal/repositories/interfaces"
    "dz-jobs-api/pkg/utils"
    "net/http"

    "github.com/google/uuid"
)

type CandidateExperienceService struct {
    candidateExperienceRepo interfaces.CandidateExperienceRepository
}

func NewCandidateExperienceService(repo interfaces.CandidateExperienceRepository) *CandidateExperienceService {
    return &CandidateExperienceService{candidateExperienceRepo: repo}
}

func (s *CandidateExperienceService) AddExperience(ctx context.Context, candidateID uuid.UUID, request request.AddExperienceRequest) (*models.CandidateExperience, error) {
    experience := &models.CandidateExperience{
        ID:          uuid.New(),
        CandidateID: candidateID,
        JobTitle:    request.JobTitle,
        Company:     request.Company,
        StartDate:   request.StartDate,
        EndDate:     request.EndDate,
        Description: request.Description,
    }

    err := s.candidateExperienceRepo.CreateExperience(ctx, experience) // Pass context
    if err != nil {
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to add experience")
    }

    return experience, nil
}

func (s *CandidateExperienceService) GetExperience(ctx context.Context, candidateID uuid.UUID) ([]models.CandidateExperience, error) {
    experiences, err := s.candidateExperienceRepo.GetExperience(ctx, candidateID) // Pass context
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, utils.NewCustomError(http.StatusNotFound, "No experience records found")
        }
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch experience records")
    }

    return experiences, nil
}

func (s *CandidateExperienceService) DeleteExperience(ctx context.Context, candidateID uuid.UUID, experienceID uuid.UUID) error {
    err := s.candidateExperienceRepo.DeleteExperience(ctx, candidateID, experienceID) // Pass context
    if err != nil {
        return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete experience")
    }

    return nil
}