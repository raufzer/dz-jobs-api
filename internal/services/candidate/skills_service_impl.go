package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	"dz-jobs-api/pkg/utils"
	models "dz-jobs-api/internal/models/candidate"
	interfaces "dz-jobs-api/internal/repositories/interfaces/candidate"
	"net/http"

	"github.com/google/uuid"
)

type CandidateSkillsService struct {
	candidateSkillsRepo interfaces.CandidateSkillsRepository
}

func NewCandidateSkillService(repo interfaces.CandidateSkillsRepository) *CandidateSkillsService {
	return &CandidateSkillsService{candidateSkillsRepo: repo}
}

func (s *CandidateSkillsService) AddSkill(candidateID uuid.UUID, request request.AddSkillRequest) (*models.CandidateSkills, error) {
	skill := &models.CandidateSkills{
		CandidateID: candidateID,
		Skill:       request.Skill,
	}

	err := s.candidateSkillsRepo.CreateSkill(skill)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to add skill")
	}

	return skill, nil
}


func (s *CandidateSkillsService) GetSkills(candidateID uuid.UUID) ([]models.CandidateSkills, error) {
	skills, err := s.candidateSkillsRepo.GetSkills(candidateID)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusNotFound, "No skills found")
	}

	return skills, nil
}

func (s *CandidateSkillsService) DeleteSkill(candidateID uuid.UUID, skill string) error {
	err := s.candidateSkillsRepo.DeleteSkill(candidateID, skill)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete skill")
	}

	return nil
}
