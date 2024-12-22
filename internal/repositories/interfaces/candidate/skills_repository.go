package candidate

import (
	models "dz-jobs-api/internal/models/candidate"
	"github.com/google/uuid"
)

type CandidateSkillsRepository interface {
	CreateSkill(skill *models.CandidateSkills) error
	GetSkills(id uuid.UUID) ([]models.CandidateSkills, error)
	DeleteSkill(candidateID uuid.UUID, skill string) error
}
