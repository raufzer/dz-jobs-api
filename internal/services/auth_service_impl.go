package services

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/helpers"
	"dz-jobs-api/internal/models"
	"dz-jobs-api/internal/repositories/interfaces"
	"dz-jobs-api/pkg/utils"

	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

type AuthService struct {
	userRepository  interfaces.UserRepository
	redisRepository interfaces.RedisRepository
}

func NewAuthService(userRepo interfaces.UserRepository, redisRepo interfaces.RedisRepository) *AuthService {
	return &AuthService{
		userRepository:  userRepo,
		redisRepository: redisRepo,
	}
}
func (s *AuthService) Register(req request.CreateUsersRequest) error {
	existingUser, _ := s.userRepository.GetByEmail(req.Email)
	if existingUser != nil {
		return helpers.NewCustomError(http.StatusBadRequest, "User already exists")
	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return helpers.NewCustomError(http.StatusInternalServerError, "Failed to hash password")
	}
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.Role,
	}
	if err := s.userRepository.Create(user); err != nil {
		return helpers.NewCustomError(http.StatusInternalServerError, "Failed to create user")
	}
	return nil
}
func (s *AuthService) Login(req request.LoginRequest) (string, string, error) {
	user, err := s.userRepository.GetByEmail(req.Email)
	if err != nil || user == nil {
		return "", "", helpers.NewCustomError(http.StatusUnauthorized, "Invalid email")
	}

	verifyErr := utils.VerifyPassword(user.Password, req.Password)
	if verifyErr != nil {
		return "", "", helpers.NewCustomError(http.StatusUnauthorized, "Invalid password")
	}

	config, err := config.LoadConfig()
	if err != nil {
		return "", "", helpers.NewCustomError(http.StatusInternalServerError, "Config loading failed")
	}

	accessToken, err := utils.GenerateToken(config.AccessTokenMaxAge, user.ID, config.AccessTokenSecret)
	if err != nil {
		return "", "", helpers.NewCustomError(http.StatusInternalServerError, "Failed to generate access token")
	}

	refreshToken, err := utils.GenerateToken(config.RefreshTokenMaxAge, user.ID, config.RefreshTokenSecret)
	if err != nil {
		return "", "", helpers.NewCustomError(http.StatusInternalServerError, "Failed to generate refresh token")
	}
	refreshTokenTTL := config.RefreshTokenMaxAge
	err = s.redisRepository.StoreRefreshToken(user.ID.String(), refreshToken, refreshTokenTTL)
	if err != nil {
		return "", "", helpers.NewCustomError(http.StatusInternalServerError, "Failed to store refresh token")
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) RefreshAccessToken(userid, refreshToken string) (string, error) {
	storedToken, err := s.redisRepository.GetRefreshToken(userid)
	if err != nil {
		if err == redis.Nil {
			return "", helpers.NewCustomError(http.StatusUnauthorized, "Refresh token expired or not found")
		}
		return "", helpers.NewCustomError(http.StatusInternalServerError, "Failed to retrieve refresh token")
	}
	if storedToken != refreshToken {
		return "", helpers.NewCustomError(http.StatusUnauthorized, "Invalid Token")
	}

	config, err := config.LoadConfig()
	if err != nil {
		return "", helpers.NewCustomError(http.StatusInternalServerError, "Config loading failed")
	}

	accessToken, err := utils.GenerateToken(config.AccessTokenMaxAge, userid, config.AccessTokenSecret)
	if err != nil {
		return "", helpers.NewCustomError(http.StatusInternalServerError, "Failed to generate access token")
	}

	return accessToken, nil
}

func (s *AuthService) SendOTP(email string) error {
	otp := utils.GenerateSecureOTP(6)
	err := s.redisRepository.StoreOTP(email, otp, 5*time.Minute)
	if err != nil {
		return helpers.NewCustomError(http.StatusInternalServerError, "Failed to store OTP")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		return helpers.NewCustomError(http.StatusInternalServerError, "Config loading failed")
	}

	return utils.SendOTPEmail(email, otp, cfg.SendGridAPIKey)
}

func (s *AuthService) VerifyOTP(email, otp string) (string, error) {
	storedOTP, err := s.redisRepository.GetOTP(email)
	if err != nil {
		if err == redis.Nil {
			return "", helpers.NewCustomError(http.StatusUnauthorized, "OTP expired or not found")
		}
		return "", helpers.NewCustomError(http.StatusInternalServerError, "Failed to retrieve OTP")
	}
	if storedOTP != otp {
		return "", helpers.NewCustomError(http.StatusUnauthorized, "Invalid OTP")
	}
	config, err := config.LoadConfig()
	if err != nil {
		return "", helpers.NewCustomError(http.StatusInternalServerError, "Config loading failed")
	}

	resetToken, err := utils.GenerateToken(config.ResetPasswordTokenMaxAge, email, config.ResetPasswordTokenSecret)
	if err != nil {
		return "", helpers.NewCustomError(http.StatusInternalServerError, "Failed to generate reset password token")
	}
	err = s.redisRepository.StoreResetToken(email, resetToken, config.ResetPasswordTokenMaxAge)
	if err != nil {
		return "", helpers.NewCustomError(http.StatusInternalServerError, "Failed to store reset token")
	}

	_ = s.redisRepository.DeleteOTP(email)

	return resetToken, nil
}

func (s *AuthService) ResetPassword(email, resetToken, newPassword string) error {
	storedToken, err := s.redisRepository.GetResetToken(email)
	if err != nil {
		if err == redis.Nil {
			return helpers.NewCustomError(http.StatusUnauthorized, "Reset token expired or not found")
		}
		return helpers.NewCustomError(http.StatusInternalServerError, "Failed to retrieve reset token")
	}
	if storedToken != resetToken {
		return helpers.NewCustomError(http.StatusUnauthorized, "Invalid reset token")
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return helpers.NewCustomError(http.StatusInternalServerError, "Failed to hash new password")
	}

	err = s.userRepository.UpdatePassword(email, string(hashedPassword))
	if err != nil {
		return helpers.NewCustomError(http.StatusInternalServerError, "Failed to update password")
	}

	_ = s.redisRepository.DeleteResetToken(email)

	return nil
}

func (s *AuthService) ValidateToken(token string) (string, error) {
	config, err := config.LoadConfig()
	if err != nil {
		return "", helpers.NewCustomError(http.StatusInternalServerError, "Config loading failed")
	}
	userID, err := utils.ValidateToken(token, config.RefreshTokenSecret)
	if err != nil {
		return "", helpers.NewCustomError(http.StatusUnauthorized, "Invalid or expired token")
	}
	return userID, nil
}
