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

type CandidateSkillsService struct {
	candidateSkillsRepo interfaces.CandidateSkillsRepository
}

func NewCandidateSkillService(repo interfaces.CandidateSkillsRepository) *CandidateSkillsService {
	return &CandidateSkillsService{candidateSkillsRepo: repo}
}

func (s *CandidateSkillsService) AddSkill(ctx context.Context, candidateID uuid.UUID, request request.AddSkillRequest) (*models.CandidateSkills, error) {
	skill := &models.CandidateSkills{
		ID:    candidateID,
		Skill: request.Skill,
	}

	err := s.candidateSkillsRepo.CreateSkill(ctx,skill)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to add skill")
	}

	return skill, nil
}

func (s *CandidateSkillsService) GetSkills(ctx context.Context, candidateID uuid.UUID) ([]models.CandidateSkills, error) {
	skills, err := s.candidateSkillsRepo.GetSkills(ctx,candidateID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "No skills found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch skills")
	}

	return skills, nil
}

func (s *CandidateSkillsService) DeleteSkill(ctx context.Context, candidateID uuid.UUID, skill string) error {
	err := s.candidateSkillsRepo.DeleteSkill(ctx,candidateID, skill)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete skill")
	}

	return nil
}
