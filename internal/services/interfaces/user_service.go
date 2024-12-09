package interfaces

import (
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"
)

type UserService interface {
	CreateUser(req request.CreateUsersRequest) (*models.User, error)
	UpdateUser(id int, req request.UpdateUserRequest) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	DeleteUser(id int) error
}
