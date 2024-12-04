package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"dz-jobs-api/data/request"
	"dz-jobs-api/data/response"
	"dz-jobs-api/internal/models"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces"
	"dz-jobs-api/pkg/helpers"
	"dz-jobs-api/pkg/utils"
)

type UserController struct {
	UserRepository repositoryInterfaces.UserRepository
}

func NewUserController(userRepository repositoryInterfaces.UserRepository) *UserController {
	return &UserController{UserRepository: userRepository}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var req request.CreateUsersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(helpers.ErrInvalidUserData)
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.Error(helpers.WrapError(err, "password hashing failed"))
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
		ctx.Error(helpers.ErrUserCreationFailed)
		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "User created successfully",
		Data:    response.ToUserResponse(user),
	})

}

func (uc *UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(errors.New("invalid user ID"))
		return
	}

	user, err := uc.UserRepository.GetByID(id)
	if err != nil || user == nil {
		ctx.Error(helpers.ErrUserNotFound)
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "User found",
		Data:    response.ToUserResponse(user),
	})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var req request.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(helpers.ErrInvalidUserData)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(errors.New("invalid user ID"))
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
		ctx.Error(helpers.WrapError(err, "failed to update user"))
		return
	}

	user, err := uc.UserRepository.GetByID(id)
	if err != nil {
		ctx.Error(helpers.WrapError(err, "failed to retrieve updated user"))
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "User updated successfully",
		Data:    response.ToUserResponse(user),
	})
}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := uc.UserRepository.GetAll()
	if err != nil {
		ctx.Error(helpers.WrapError(err, "failed to fetch users"))
		return
	}

	userResponses := make([]response.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = response.ToUserResponse(user)
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Users retrieved successfully",
		Data: gin.H{
			"total": len(users),
			"users": userResponses,
		},
	})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(errors.New("invalid user ID"))
		return
	}

	if err := uc.UserRepository.Delete(id); err != nil {
		ctx.Error(helpers.ErrUserNotFound)
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "User deleted successfully",
	})
}
