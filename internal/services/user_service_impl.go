package services

import (
	"database/sql"
	"dz-jobs-api/internal/dto/request"
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

func (s *UserService) CreateUser(req request.CreateUsersRequest) (*models.User, error) {
	existingUser, err := s.userRepository.GetUserByEmail(req.Email)
	if err != nil && err != sql.ErrNoRows {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Database error occurred")
	}
	if existingUser != nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "User already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Password hashing failed")
	}

	user := &models.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		Role:      req.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.userRepository.CreateUser(user); err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "User creation failed")
	}

	return user, nil
}

func (s *UserService) GetUser(id uuid.UUID) (*models.User, error) {
	user, err := s.userRepository.GetUserByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "User not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching user")
	}
	return user, nil
}

func (s *UserService) UpdateUser(id uuid.UUID, req request.UpdateUserRequest) (*models.User, error) {
	updatedUser := &models.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		Role:      req.Role,
		UpdatedAt: time.Now(),
	}

	if err := s.userRepository.UpdateUser(id, updatedUser); err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "User not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to update user")
	}

	return s.userRepository.GetUserByID(id)
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch users")
	}
	return users, nil
}

func (s *UserService) DeleteUser(id uuid.UUID) error {
	err := s.userRepository.DeleteUser(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.NewCustomError(http.StatusNotFound, "User not found")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete user")
	}
	return nil
}
