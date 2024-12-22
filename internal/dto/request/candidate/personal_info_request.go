package candidate


type AddPersonalInfoRequest struct {
	Name        string    `json:"name" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
}
type UpdatePersonalInfoRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
