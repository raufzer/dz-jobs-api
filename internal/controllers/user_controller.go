package controllers

import (
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/dto/response"
	serviceInterfaces "dz-jobs-api/internal/services/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userService serviceInterfaces.UserService
}

func NewUserController(service serviceInterfaces.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags Admin - Users
// @Accept json
// @Produce json
// @Param user body request.CreateUsersRequest true "User request"
// @Success 201 {object} response.Response{Data=response.UserResponse} "User created successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 400 {object} response.Response "Invalid user ID"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "User not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /admin/users [post]
func (c *UserController) CreateUser(ctx *gin.Context) {
	var req request.CreateUsersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_  = ctx.Error(err)
		ctx.Abort()
		return
	}
	user, err := c.userService.CreateUser(ctx,req)
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "User created successfully",
		Data:    response.ToUserResponse(user),
	})
}

// GetUser godoc
// @Summary Get user
// @Description Get user details by user ID
// @Tags Admin - Users
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} response.Response{Data=response.UserResponse} "User found"
// @Failure 400 {object} response.Response "Invalid user ID"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "User not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /admin/users/{userId} [get]
func (c *UserController) GetUser(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		_  = ctx.Error(err)
		ctx.Abort()
		return
	}
	user, err := c.userService.GetUser(ctx,userID)
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "User found",
		Data:    response.ToUserResponse(user),
	})
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user details by user ID
// @Tags Admin - Users
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param user body request.UpdateUserRequest true "User request"
// @Success 200 {object} response.Response{Data=response.UserResponse} "User updated successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 400 {object} response.Response "Invalid user ID"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "User not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /admin/users/{userId} [patch]
func (c *UserController) UpdateUser(ctx *gin.Context) {
	var req request.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_  = ctx.Error(err)
		ctx.Abort()
		return
	}
	id, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	updatedUser, err := c.userService.UpdateUser(ctx,id, req)
	if err != nil {
		_  = ctx.Error(err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "User updated successfully",
		Data:    response.ToUserResponse(updatedUser),
	})
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags Admin - Users
// @Produce json
// @Success 200 {object} response.Response{Data=response.UsersResponseData} "Users retrieved successfully"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /admin/users [get]
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers(ctx)
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Users retrieved successfully",
		Data:    response.ToUsersResponse(users),
	})
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user by user ID
// @Tags Admin - Users
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} response.Response"User deleted successfully"
// @Failure 400 {object} response.Response "Invalid user ID"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "User not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /admin/users/{userId} [delete]
func (c *UserController) DeleteUser(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		_  = ctx.Error(err)
		ctx.Abort()
		return
	}
	err = c.userService.DeleteUser(ctx,userID)
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "User deleted successfully",
	})
}
