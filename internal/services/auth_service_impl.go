package services

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/helpers"
	"dz-jobs-api/internal/models"
	"dz-jobs-api/internal/repositories/interfaces"
	"dz-jobs-api/pkg/utils"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"net/http"
	"time"
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
func (s *AuthService) Login(req request.LoginRequest) (string, error) {
	user, err := s.userRepository.GetByEmail(req.Email)
	if err != nil || user == nil {
		return "", helpers.NewCustomError(http.StatusUnauthorized, "Invalid email")
	}
	config, err := config.LoadConfig()
	if err != nil {
		return "", helpers.NewCustomError(http.StatusInternalServerError, "Config loading failed")
	}
	verifyErr := utils.VerifyPassword(user.Password, req.Password)
	if verifyErr != nil {
		return "", helpers.NewCustomError(http.StatusUnauthorized, "Invalid password")
	}
	token, err := utils.GenerateToken(config.TokenExpiresIn, user.ID, config.TokenSecret)
	if err != nil {
		return "", helpers.NewCustomError(http.StatusInternalServerError, "Token generation failed")
	}
	return token, nil
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
func (s *AuthService) SendOTP(email string) error {
	otp := utils.GenerateSecureOTP(6)
	err := s.redisRepository.StoreOTP(email, otp, 5*time.Minute)
	if err != nil {
		return helpers.NewCustomError(http.StatusInternalServerError, "Failed to store OTP")
	}
	cfg, err := config.LoadConfig()
	return utils.SendOTP(email, otp, cfg.SendGridAPIKey)
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
	resetToken := uuid.NewString()
	err = s.redisRepository.StoreResetToken(email, resetToken, 5*time.Minute)
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
