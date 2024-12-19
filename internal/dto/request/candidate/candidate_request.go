package candidate
type CreateCandidateRequest struct {
	Resume         string `json:"resume" binding:"required"`
	ProfilePicture string `json:"profile_picture"`
}

type UpdateCandidateRequest struct {
	Resume         string `json:"resume"`
	ProfilePicture string `json:"profile_picture"`
}
