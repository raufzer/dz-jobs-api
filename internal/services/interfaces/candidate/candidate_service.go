package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	models "dz-jobs-api/internal/models/candidate"

	"github.com/google/uuid"
)

type CandidateService interface {
	CreateCandidate(request request.CreateCandidateRequest) (*models.Candidate, error)
	GetCandidateByID(candidateID uuid.UUID) (*models.Candidate, error)
	UpdateCandidate(candidateID uuid.UUID, req request.UpdateCandidateRequest) (*models.Candidate ,error)
	DeleteCandidate(candidateID uuid.UUID) error
}
