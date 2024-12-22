package interfaces

import (
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(user_id uuid.UUID) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(user_id uuid.UUID, user *models.User) error
	UpdateUserPassword(email, hashedPassword string) error
	DeleteUser(user_id uuid.UUID) error
}
