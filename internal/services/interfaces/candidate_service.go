package interfaces

import (
	"dz-jobs-api/internal/models"
	"mime/multipart"

	"github.com/google/uuid"
)

type CandidateService interface {
	CreateCandidate(userID string, profilePictureFile, resumeFile *multipart.FileHeader) (*models.Candidate, error)
	GetCandidate(candidateID uuid.UUID) (*models.Candidate, error)
	UpdateCandidate(candidateID uuid.UUID, profilePictureFile, resumeFile *multipart.FileHeader) (*models.Candidate, error)
	DeleteCandidate(candidateID uuid.UUID) error
}
