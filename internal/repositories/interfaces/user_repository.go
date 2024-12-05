package interfaces

import "dz-jobs-api/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByID(userid int) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(userid int, user *models.User) error
	Delete(userid int) error
}
