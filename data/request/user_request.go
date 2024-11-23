package request

// CreateUsersRequest represents the request structure for user registration
type CreateUsersRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role,omitempty"` 
}

// LoginRequest represents the request structure for user login
type LoginRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// UpdateUserRequest represents the request structure for updating user information
type UpdateUserRequest struct {
	Name     string `json:"name,omitempty" validate:"omitempty,min=3,max=50"`
	Email    string `json:"email,omitempty" validate:"omitempty,email"`
	Password string `json:"password,omitempty" validate:"omitempty,min=6"`
	Role     string `json:"role,omitempty"`
}
