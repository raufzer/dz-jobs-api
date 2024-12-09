package services

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/helpers"
	"dz-jobs-api/internal/models"
	"dz-jobs-api/internal/repositories/interfaces"
	"dz-jobs-api/pkg/utils"
	"net/http"
)

type AuthService struct {
	userRepository interfaces.UserRepository
}

func NewAuthService(userRepo interfaces.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepo}
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
