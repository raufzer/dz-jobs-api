package candidate

import (
	models "dz-jobs-api/internal/models/candidate"
	"mime/multipart"

	"github.com/google/uuid"
)

type CandidateService interface {
	CreateCandidate(profilePictureFile, resumeFile *multipart.FileHeader) (*models.Candidate, error)
	GetCandidateByID(candidateID uuid.UUID) (*models.Candidate, error)
	UpdateCandidate(candidateID uuid.UUID, profilePictureFile, resumeFile *multipart.FileHeader) (*models.Candidate, error)
	DeleteCandidate(candidateID uuid.UUID) error
}
