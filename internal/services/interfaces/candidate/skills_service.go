package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	models "dz-jobs-api/internal/models/candidate"
	"github.com/google/uuid"
)

type CandidateSkillsService interface {
	AddSkill(candidateID uuid.UUID, request request.AddSkillRequest) (*models.CandidateSkills, error)
	GetSkillsByCandidateID(candidateID uuid.UUID) ([]models.CandidateSkills, error)
	DeleteSkill(candidateID uuid.UUID, skill string) error
}
