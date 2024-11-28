package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"dz-jobs-api/data/request"
	"dz-jobs-api/data/response"
	"dz-jobs-api/internal/models"
	"dz-jobs-api/internal/repositories"
	"dz-jobs-api/pkg/helpers"
	"dz-jobs-api/pkg/utils"
)

type UserController struct {
	UserRepository repositories.UserRepository
}

func NewUserController(userRepository repositories.UserRepository) *UserController {
	return &UserController{UserRepository: userRepository}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var req request.CreateUsersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.RespondWithError(ctx, "Invalid request format", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		helpers.RespondWithError(ctx, "Password hashing failed", http.StatusInternalServerError)
		return
	}

	user := &models.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.UserRepository.Create(user); err != nil {
		helpers.RespondWithError(ctx, "Failed to create user", http.StatusInternalServerError)
		return
	}

	helpers.RespondWithSuccess(ctx, "User created successfully", toUserResponse(user))
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helpers.RespondWithError(ctx, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := uc.UserRepository.GetByID(id)
	if err != nil || user == nil {
		helpers.RespondWithError(ctx, "User not found", http.StatusNotFound)
		return
	}

	helpers.RespondWithSuccess(ctx, "User found", toUserResponse(user))
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var req request.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.RespondWithError(ctx, "Invalid request format", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helpers.RespondWithError(ctx, "Invalid user ID", http.StatusBadRequest)
		return
	}

	updatedUser := &models.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		Role:      req.Role,
		UpdatedAt: time.Now(),
	}

	if err := uc.UserRepository.Update(id, updatedUser); err != nil {
		helpers.RespondWithError(ctx, "Failed to update user", http.StatusInternalServerError)
		return
	}

	user, err := uc.UserRepository.GetByID(id)
	if err != nil {
		helpers.RespondWithError(ctx, "Failed to retrieve updated user", http.StatusInternalServerError)
		return
	}

	helpers.RespondWithSuccess(ctx, "User updated successfully", toUserResponse(user))
}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := uc.UserRepository.GetAll()
	if err != nil {
		helpers.RespondWithError(ctx, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	userResponses := make([]response.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = toUserResponse(user)
	}

	helpers.RespondWithSuccess(ctx, "Users retrieved successfully", userResponses)
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helpers.RespondWithError(ctx, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := uc.UserRepository.Delete(id); err != nil {
		helpers.RespondWithError(ctx, "User not found", http.StatusNotFound)
		return
	}

	helpers.RespondWithSuccess(ctx, "User deleted successfully", nil)
}

func toUserResponse(user *models.User) response.UserResponse {
	return response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
