package interfaces

import (
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(req request.CreateUsersRequest) (*models.User, error)
	UpdateUser(userID uuid.UUID, req request.UpdateUserRequest) (*models.User, error)
	GetUser(userID uuid.UUID) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	DeleteUser(userID uuid.UUID) error
}
