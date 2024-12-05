package request

type CreateCandidateRequest struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required,min=8"`
	Role           string `json:"role" binding:"required,min=4"`
	Resume         string `json:"resume" binding:"required"`
	Portfolio      string `json:"portfolio" binding:"required"`
	Skills         string `json:"skills" binding:"required"`
	TestID         int    `json:"testId" binding:"required"`
	ProfilePicture string `json:"profilePicture" binding:"required"`
}

type UpdateCandidateRequest struct {
	Name           string `json:"name,omitempty" validate:"omitempty,min=3,max=50"`
	Email          string `json:"email,omitempty" validate:"omitempty,email"`
	Password       string `json:"password,omitempty" validate:"omitempty,min=6"`
	Role           string `json:"role,omitempty" validate:"omitempty,min=4"`
	Resume         string `json:"resume,omitempty"`
	Portfolio      string `json:"portfolio,omitempty"`
	Skills         string `json:"skills,omitempty"`
	TestID         int    `json:"testId,omitempty"`
	ProfilePicture string `json:"profilePicture,omitempty"`
}
