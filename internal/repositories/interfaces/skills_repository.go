package interfaces

import (
	"dz-jobs-api/internal/models"
	"github.com/google/uuid"
)

type CandidateSkillsRepository interface {
	CreateSkill(skill *models.CandidateSkills) error
	GetSkills(candidateID uuid.UUID) ([]models.CandidateSkills, error)
	DeleteSkill(candidateID uuid.UUID, skill string) error
}
