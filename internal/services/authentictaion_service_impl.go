package services

import (
	"dz-jobs-api/config"
	"dz-jobs-api/data/request"
	"dz-jobs-api/internal/models"
	"dz-jobs-api/internal/repositories"
	"dz-jobs-api/pkg/helpers"
	"dz-jobs-api/pkg/utils"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type AuthenticationServiceImpl struct {
	UsersRepository repositories.UserRepository
	Validate        *validator.Validate
}

func NewAuthenticationServiceImpl(usersRepository repositories.UserRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UsersRepository: usersRepository,
		Validate:        validate,
	}
}

// Login implements AuthenticationService
func (a *AuthenticationServiceImpl) Login(user request.LoginRequest) (string, error) {
	// Find username in database
	new_users, user_err := a.UsersRepository.GetByName(user.Name)
	if user_err != nil {
		return "", errors.New("invalid username or Password")
	}

	config, err := config.LoadConfig()
	helpers.ErrorPanic(err) // Panic on configuration error as it's critical

	verify_error := utils.VerifyPassword(new_users.Password, user.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or Password")
	}

	// Generate Token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_users.ID, config.TokenSecret)
	helpers.ErrorPanic(err_token) // Panic on token generation error as it's a critical system function

	return token, nil
}

// Register implements AuthenticationService
func (a *AuthenticationServiceImpl) Register(user request.CreateUsersRequest) {
	// Validate user input
	err := a.Validate.Struct(user)
	helpers.ErrorPanic(err) // Panic on validation error as it indicates a system issue

	// Check if user already exists
	_, err = a.UsersRepository.GetByName(user.Name)
	if err == nil {
		helpers.ErrorPanic(errors.New("username already exists"))
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	helpers.ErrorPanic(err) // Panic on hashing error as it's a critical security function

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
	err = a.UsersRepository.Create(&newUser)
	helpers.ErrorPanic(err) // Panic on database error as it's a critical system function
}
