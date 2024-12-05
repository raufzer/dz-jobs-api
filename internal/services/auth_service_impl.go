package services

import (
	"dz-jobs-api/config"
	"dz-jobs-api/helpers"
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces"
	serviceInterfaces "dz-jobs-api/internal/services/interfaces"
	"dz-jobs-api/pkg/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	UserRepository repositoryInterfaces.UserRepository
	Validate       *validator.Validate
}

func NewAuthServiceImpl(userRepository repositoryInterfaces.UserRepository, validate *validator.Validate) serviceInterfaces.AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (a *AuthServiceImpl) Login(user request.LoginRequest) (string, error) {
	newUser, userErr := a.UserRepository.GetByEmail(user.Email)
	if userErr != nil {
		return "", helpers.ErrInvalidCredentials
	}
	config, err := config.LoadConfig()
	if err != nil {
		return "", helpers.WrapError(err, "config loading failed")
	}

	verifyErr := utils.VerifyPassword(newUser.Password, user.Password)
	if verifyErr != nil {
		return "", helpers.ErrInvalidCredentials
	}

	token, errToken := utils.GenerateToken(config.TokenExpiresIn, newUser.ID, config.TokenSecret)
	if errToken != nil {
		return "", helpers.ErrTokenGeneration
	}

	return token, nil
}

func (a *AuthServiceImpl) Register(user request.CreateUsersRequest) error {

	if err := a.Validate.Struct(user); err != nil {
		return helpers.WrapError(err, "validation failed")
	}

	existingUser, err := a.UserRepository.GetByEmail(user.Email)
	if err == nil && existingUser != nil {
		return helpers.ErrEmailAlreadyExists
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return helpers.WrapError(err, "password hashing failed")
	}

	now := time.Now()
	newUser := models.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  hashedPassword,
		Role:      user.Role,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := a.UserRepository.Create(&newUser); err != nil {
		return helpers.WrapError(err, "user creation failed")
	}

	return nil
}
