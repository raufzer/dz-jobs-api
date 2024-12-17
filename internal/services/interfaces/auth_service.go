package interfaces

import (
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"
)

type AuthService interface {
	Register(user request.CreateUsersRequest) (*models.User, error)
	Login(req request.LoginRequest) (*models.User, string, string, error)
	RefreshAccessToken(email, role, refreshToken string) (string, error)
	SendOTP(email string) error
	VerifyOTP(email, otp string) (string, error)
	ResetPassword(email, resetToken, newPassword string) error
	ValidateToken(token string) (string, string, error)
	GoogleConnect(code string) (*models.User, string, string, string, error)
}
