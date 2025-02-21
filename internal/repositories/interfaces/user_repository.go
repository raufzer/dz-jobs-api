package interfaces

import (
	"context"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	GetAllUsers(ctx context.Context, ) ([]*models.User, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, user *models.User) error
	UpdateUserPassword(ctx context.Context, email, hashedPassword string) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}
