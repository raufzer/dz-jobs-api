package interfaces

import "dz-jobs-api/internal/dto/request"

type AuthService interface {
	Login(user request.LoginRequest) (string, error)
	Register(user request.CreateUsersRequest) error
	SendOTP(email string) error
	VerifyOTP(email, otp string) (string, error)
	ResetPassword(email, resetToken, newPassword string) error
}
