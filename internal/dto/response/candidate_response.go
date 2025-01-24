package response

import (
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type CandidateResponse struct {
	ID             uuid.UUID `json:"candidate_id"`
	Resume         string    `json:"resume"`
	ProfilePicture string    `json:"profile_picture"`
}

func ToCandidateResponse(candidate *models.Candidate) CandidateResponse {
	return CandidateResponse{
		ID:             candidate.ID,
		Resume:         candidate.Resume,
		ProfilePicture: candidate.ProfilePicture,
	}
}
