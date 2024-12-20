package candidate


type CreateCandidatePersonalInfoRequest struct {
	Name        string    `json:"name" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
}
type UpdateCandidatePersonalInfoRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
