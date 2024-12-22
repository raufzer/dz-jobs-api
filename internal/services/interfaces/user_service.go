package interfaces

import (
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(req request.CreateUsersRequest) (*models.User, error)
	UpdateUser(id uuid.UUID, req request.UpdateUserRequest) (*models.User, error)
	GetUser(id uuid.UUID) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	DeleteUser(id uuid.UUID) error
}
