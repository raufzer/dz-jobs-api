package services

import (
	"database/sql"
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/helpers"
	"dz-jobs-api/internal/models"
	"dz-jobs-api/internal/repositories/interfaces"
	"dz-jobs-api/pkg/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(userRepo interfaces.UserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

func (us *UserService) CreateUser(req request.CreateUsersRequest) (*models.User, error) {
	existingUser, err := us.userRepository.GetByEmail(req.Email)
	if err != nil && err != sql.ErrNoRows {
		return nil, helpers.NewCustomError(http.StatusInternalServerError, "Database error occurred")
	}
	if existingUser != nil {
		return nil, helpers.NewCustomError(http.StatusBadRequest, "User already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, helpers.NewCustomError(http.StatusInternalServerError, "Password hashing failed")
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.Role,
	}

	if err := us.userRepository.Create(user); err != nil {
		return nil, helpers.NewCustomError(http.StatusInternalServerError, "User creation failed")
	}

	return user, nil
}

func (us *UserService) GetUserByID(id uuid.UUID) (*models.User, error) {
	user, err := us.userRepository.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helpers.NewCustomError(http.StatusNotFound, "User not found")
		}
		return nil, helpers.NewCustomError(http.StatusInternalServerError, "Error fetching user")
	}
	return user, nil
}

func (us *UserService) UpdateUser(id uuid.UUID, req request.UpdateUserRequest) (*models.User, error) {
	updatedUser := &models.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		Role:      req.Role,
		UpdatedAt: time.Now(),
	}

	if err := us.userRepository.Update(id, updatedUser); err != nil {
		if err == sql.ErrNoRows {
			return nil, helpers.NewCustomError(http.StatusNotFound, "User not found")
		}
		return nil, helpers.NewCustomError(http.StatusInternalServerError, "Failed to update user")
	}

	return us.userRepository.GetByID(id)
}

func (us *UserService) GetAllUsers() ([]*models.User, error) {
	users, err := us.userRepository.GetAll()
	if err != nil {
		return nil, helpers.NewCustomError(http.StatusInternalServerError, "Failed to fetch users")
	}
	return users, nil
}

func (us *UserService) DeleteUser(id uuid.UUID) error {
	err := us.userRepository.Delete(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return helpers.NewCustomError(http.StatusNotFound, "User not found")
		}
		return helpers.NewCustomError(http.StatusInternalServerError, "Failed to delete user")
	}
	return nil
}
