package request

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type SendOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type VerifyOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required,len=6"`
}
type ResetPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
	// ResetToken  string `json:"reset_token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}
