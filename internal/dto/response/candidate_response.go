package response

import (
	"dz-jobs-api/internal/models"
)

type CandidateResponse struct {
	User           UserResponse `json:"user"`
	CandidateID    int          `json:"candidateId"`
	Resume         string       `json:"resume"`
	Portfolio      string       `json:"portfolio"`
	Skills         string       `json:"skills"`
	TestID         int          `json:"testId"`
	ProfilePicture string       `json:"profilePicture"`

}

func ToCandidateResponse(candidate *models.Candidate) CandidateResponse {
	return CandidateResponse{
		User:          ToUserResponse(&candidate.User),
		CandidateID:   candidate.CandidateID,
		Resume:        candidate.Resume,
		Portfolio:     candidate.Portfolio,
		Skills:        candidate.Skills,
		TestID:        candidate.TestID,
		ProfilePicture: candidate.ProfilePicture,
	}
}