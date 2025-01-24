package interfaces

import (
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(userID uuid.UUID) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(userID uuid.UUID, user *models.User) error
	UpdateUserPassword(email, hashedPassword string) error
	DeleteUser(userID uuid.UUID) error
}
