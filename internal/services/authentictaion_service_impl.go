package services

import (
	"dz-jobs-api/config"
	"dz-jobs-api/data/request"
	"dz-jobs-api/internal/models"
	"dz-jobs-api/internal/repositories"
	"dz-jobs-api/pkg/utils"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type AuthenticationServiceImpl struct {
	UserRepository repositories.UserRepository
	Validate        *validator.Validate
}

// NewAuthenticationServiceImpl returns a pointer to AuthenticationServiceImpl
func NewAuthenticationServiceImpl(userRepository repositories.UserRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{  // Return a pointer to AuthenticationServiceImpl
		UserRepository: userRepository,
		Validate:        validate,
	}
}

// Login implements AuthenticationService
func (a *AuthenticationServiceImpl) Login(user request.LoginRequest) (string, error) {
	// Find username in database
	newUser, userErr := a.UserRepository.GetByName(user.Name)
	if userErr != nil {
		return "", errors.New("invalid name or password") // Return specific error message
	}

	// Load configuration
	config, err := config.LoadConfig()
	if err != nil {
		return "", err // Return error instead of panicking
	}

	// Verify password
	verifyErr := utils.VerifyPassword(newUser.Password, user.Password)
	if verifyErr != nil {
		return "", errors.New("invalid name or password") // Return specific error message
	}

	// Generate Token
	token, errToken := utils.GenerateToken(config.TokenExpiresIn, newUser.ID, config.TokenSecret)
	if errToken != nil {
		return "", errToken // Return error instead of panicking
	}

	return token, nil
}

// Register implements AuthenticationService
func (a *AuthenticationServiceImpl) Register(user request.CreateUsersRequest) error {
	// Validate user input
	err := a.Validate.Struct(user)
	if err != nil {
		return err // Return validation error instead of panicking
	}

	// Check if user already exists
	_, err = a.UserRepository.GetByName(user.Name)
	if err == nil {
		return errors.New("username already exists") // Return error instead of panicking
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err // Return hashing error instead of panicking
	}

	// Create new user with all required fields
	now := time.Now()
	newUser := models.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  hashedPassword,
		Role:      "user", // Set default role
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Create user in repository
	err = a.UserRepository.Create(&newUser)
	if err != nil {
		return err // Return database error instead of panicking
	}

	return nil // Return nil if registration is successful
}
