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

type AuthServiceImpl struct {
	UserRepository repositories.UserRepository
	Validate       *validator.Validate
}

func NewAuthServiceImpl(userRepository repositories.UserRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (a *AuthServiceImpl) Login(user request.LoginRequest) (string, error) {

	newUser, userErr := a.UserRepository.GetByName(user.Email)
	if userErr != nil {
		return "", errors.New("invalid name or password")
	}

	config, err := config.LoadConfig()
	if err != nil {
		return "", err
	}

	verifyErr := utils.VerifyPassword(newUser.Password, user.Password)
	if verifyErr != nil {
		return "", errors.New("invalid name or password")
	}

	token, errToken := utils.GenerateToken(config.TokenExpiresIn, newUser.ID, config.TokenSecret)
	if errToken != nil {
		return "", errToken
	}

	return token, nil
}

func (a *AuthServiceImpl) Register(user request.CreateUsersRequest) error {

	err := a.Validate.Struct(user)
	if err != nil {
		return err
	}

	_, err = a.UserRepository.GetByName(user.Name)
	if err == nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	now := time.Now()
	newUser := models.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  hashedPassword,
		Role:      "user",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = a.UserRepository.Create(&newUser)
	if err != nil {
		return err
	}

	return nil
}
