package interfaces

import (
    "context"
    "dz-jobs-api/internal/dto/request"
    "dz-jobs-api/internal/models"

    "github.com/google/uuid"
)

type UserService interface {
    CreateUser(ctx context.Context, req request.CreateUsersRequest) (*models.User, error)
    UpdateUser(ctx context.Context, userID uuid.UUID, req request.UpdateUserRequest) (*models.User, error)
    GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error)
    GetAllUsers(ctx context.Context) ([]*models.User, error)
    DeleteUser(ctx context.Context, userID uuid.UUID) error
}