package interfaces

import "dz-jobs-api/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByID(id int) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(id int, user *models.User) error
	Delete(id int) error
}
