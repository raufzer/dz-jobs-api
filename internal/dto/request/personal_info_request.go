package request

type AddPersonalInfoRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender" binding:"required,oneof=male female"`
	Bio         string `json:"bio"`
}

type UpdatePersonalInfoRequest struct {
	Name        string `json:"name" binding:"omitempty"`
	Email       string `json:"email" binding:"omitempty,email"`
	Phone       string `json:"phone" binding:"omitempty"`
	Address     string `json:"address" binding:"omitempty"`
	DateOfBirth string `json:"date_of_birth" binding:"omitempty"`
	Gender      string `json:"gender" binding:"omitempty,oneof=male female"`
	Bio         string `json:"bio" binding:"omitempty"`
}
