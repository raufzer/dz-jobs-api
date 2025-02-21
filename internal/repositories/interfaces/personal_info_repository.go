package interfaces

import (
	"context"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type CandidatePersonalInfoRepository interface {
	CreatePersonalInfo(ctx context.Context, info *models.CandidatePersonalInfo) error
	GetPersonalInfo(ctx context.Context, candidateID uuid.UUID) (*models.CandidatePersonalInfo, error)
	UpdatePersonalInfo(ctx context.Context, info *models.CandidatePersonalInfo) error
	DeletePersonalInfo(ctx context.Context, candidateID uuid.UUID) error
}
