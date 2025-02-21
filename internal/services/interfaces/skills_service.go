package interfaces

import (
    "context"
    "dz-jobs-api/internal/dto/request"
    "dz-jobs-api/internal/models"
    "github.com/google/uuid"
)

type CandidateSkillsService interface {
    AddSkill(ctx context.Context, candidateID uuid.UUID, request request.AddSkillRequest) (*models.CandidateSkills, error)
    GetSkills(ctx context.Context, candidateID uuid.UUID) ([]models.CandidateSkills, error)
    DeleteSkill(ctx context.Context, candidateID uuid.UUID, skill string) error
}