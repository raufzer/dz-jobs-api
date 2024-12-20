package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	models "dz-jobs-api/internal/models/candidate"
	"github.com/google/uuid"
)

type CandidatePersonalInfoService interface {
	UpdatePersonalInfo(id uuid.UUID,request request.UpdateCandidatePersonalInfoRequest) (*models.CandidatePersonalInfo, error)
	GetPersonalInfo(candidateID uuid.UUID) (*models.CandidatePersonalInfo, error)
}
