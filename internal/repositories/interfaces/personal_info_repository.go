package interfaces

import (
	"dz-jobs-api/internal/models"
	"github.com/google/uuid"
)

type CandidatePersonalInfoRepository interface {
	CreatePersonalInfo(info *models.CandidatePersonalInfo) error
	GetPersonalInfo(id uuid.UUID) (*models.CandidatePersonalInfo, error)
	UpdatePersonalInfo(info *models.CandidatePersonalInfo) error
	DeletePersonalInfo(id uuid.UUID) error
}
