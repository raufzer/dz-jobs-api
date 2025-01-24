package response

import (
	"dz-jobs-api/internal/models"
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToUserResponse(user *models.User) UserResponse {
	return UserResponse{
		UserID:    user.UserID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type UsersResponseData struct {
	Total int            `json:"total"`
	Users []UserResponse `json:"users"`
}

func ToUsersResponse(users []*models.User) UsersResponseData {
	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return UsersResponseData{
		Total: len(users),
		Users: userResponses,
	}
}
