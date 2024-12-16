package interfaces

import (
	"dz-jobs-api/internal/dto/request"
)

type AuthService interface {
	Register(user request.CreateUsersRequest) error
	Login(req request.LoginRequest) (string, string, error)
	RefreshAccessToken(email, refreshToken string) (string, error)
	SendOTP(email string) error
	VerifyOTP(email, otp string) (string, error)
	ResetPassword(email, resetToken, newPassword string) error
	ValidateToken(token string) (string, error)
}
