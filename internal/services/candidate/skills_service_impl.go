package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	"dz-jobs-api/pkg/utils"
	models "dz-jobs-api/internal/models/candidate"
	interfaces "dz-jobs-api/internal/repositories/interfaces/candidate"
	"net/http"

	"github.com/google/uuid"
)

type candidateSkillsService struct {
	skillsRepo interfaces.CandidateSkillsRepository
}

func NewCandidateSkillService(repo interfaces.CandidateSkillsRepository) *candidateSkillsService {
	return &candidateSkillsService{skillsRepo: repo}
}

func (s *candidateSkillsService) AddSkill(candidateID uuid.UUID, request request.AddSkillRequest) (*models.CandidateSkills, error) {
	skill := &models.CandidateSkills{
		CandidateID: candidateID,
		Skill:       request.Skill,
	}

	err := s.skillsRepo.CreateSkill(skill)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to add skill")
	}

	return skill, nil
}

func (s *candidateSkillsService) GetSkillsByCandidateID(candidateID uuid.UUID) ([]models.CandidateSkills, error) {
	skills, err := s.skillsRepo.GetSkillsByCandidateID(candidateID)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusNotFound, "No skills found")
	}

	return skills, nil
}

func (s *candidateSkillsService) DeleteSkill(candidateID uuid.UUID, skill string) error {
	err := s.skillsRepo.DeleteSkill(candidateID, skill)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete skill")
	}

	return nil
}
