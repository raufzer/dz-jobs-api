package candidate

import (
	models "dz-jobs-api/internal/models/candidate"
	"github.com/google/uuid"
)

type CandidateRepository interface {
	CreateCandidate(candidate models.Candidate) (uuid.UUID, error)
	GetCandidateByID(id uuid.UUID) (models.Candidate, error)
	UpdateCandidate(candidate_id uuid.UUID, candidate models.Candidate) error
	DeleteCandidate(id uuid.UUID) error
}
