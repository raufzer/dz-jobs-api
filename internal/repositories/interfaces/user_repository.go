package interfaces

import (
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByID(userid uuid.UUID) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(userid uuid.UUID, user *models.User) error
	UpdatePassword(email, hashedPassword string) error
	Delete(userid uuid.UUID) error
}
