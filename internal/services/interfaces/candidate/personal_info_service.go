package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	models "dz-jobs-api/internal/models/candidate"
	"github.com/google/uuid"
)

type CandidatePersonalInfoService interface {
	AddPersonalInfo(request request.AddPersonalInfoRequest, candidateID uuid.UUID) (*models.CandidatePersonalInfo, error)
	UpdatePersonalInfo(id uuid.UUID, request request.UpdatePersonalInfoRequest) (*models.CandidatePersonalInfo, error)
	GetPersonalInfo(candidateID uuid.UUID) (*models.CandidatePersonalInfo, error)
	DeletePersonalInfo(candidateID uuid.UUID) error
}
