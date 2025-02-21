package interfaces

import (
	"context"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type CandidateSkillsRepository interface {
	CreateSkill(ctx context.Context, skill *models.CandidateSkills) error
	GetSkills(ctx context.Context, candidateID uuid.UUID) ([]models.CandidateSkills, error)
	DeleteSkill(ctx context.Context, candidateID uuid.UUID, skill string) error
}
