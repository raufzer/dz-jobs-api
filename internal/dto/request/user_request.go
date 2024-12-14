package request

type CreateUsersRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Role     string `json:"role" binding:"required,min=4"`
}

type UpdateUserRequest struct {
	Name     string `json:"name,omitempty" validate:"omitempty,min=3,max=50"`
	Email    string `json:"email,omitempty" validate:"omitempty,email"`
	Password string `json:"password,omitempty" validate:"omitempty,min=6"`
	Role     string `json:"role,omitempty" validate:"omitempty,min=4"`
}
