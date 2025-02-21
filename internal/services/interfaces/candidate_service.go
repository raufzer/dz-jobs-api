package interfaces

import (
    "context"
    "dz-jobs-api/internal/models"
    "mime/multipart"

    "github.com/google/uuid"
)

type CandidateService interface {
    CreateCandidate(ctx context.Context, userID string, profilePictureFile, resumeFile *multipart.FileHeader) (*models.Candidate, error)
    CreateDefaultCandidate(ctx context.Context, userID, profilePictureDefault, resumeDefault string) (*models.Candidate, error)
    GetCandidate(ctx context.Context, candidateID uuid.UUID) (*models.Candidate, error)
    UpdateCandidate(ctx context.Context, candidateID uuid.UUID, profilePictureFile, resumeFile *multipart.FileHeader) (*models.Candidate, error)
    DeleteCandidate(ctx context.Context, candidateID uuid.UUID) error
}