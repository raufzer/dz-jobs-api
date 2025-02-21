package interfaces

import (
    "context"
    "dz-jobs-api/internal/dto/request"
    "dz-jobs-api/internal/models"
)

type AuthService interface {
    Register(ctx context.Context, user request.CreateUsersRequest) (*models.User, error)
    Login(ctx context.Context, req request.LoginRequest) (*models.User, string, string, error)
    Logout(ctx context.Context, userID, refreshToken string) error
    RefreshAccessToken(ctx context.Context, email, role, refreshToken string) (string, error)
    SendOTP(ctx context.Context, email string) error
    VerifyOTP(ctx context.Context, email, otp string) (string, error)
    ResetPassword(ctx context.Context, email, resetToken, newPassword string) error
    ValidateToken(ctx context.Context, token string) (string, string, error)
    GoogleConnect(ctx context.Context, code string, role string) (*models.User, string, string, string, error)
}