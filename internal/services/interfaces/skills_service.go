package interfaces

import (
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"
	"github.com/google/uuid"
)

type CandidateSkillsService interface {
	AddSkill(candidateID uuid.UUID, request request.AddSkillRequest) (*models.CandidateSkills, error)
	GetSkills(candidateID uuid.UUID) ([]models.CandidateSkills, error)
	DeleteSkill(candidateID uuid.UUID, skill string) error
}
