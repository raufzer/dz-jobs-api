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
// @Tags users_1create
// @Accept json
// @Produce json
// @Param user body request.CreateUsersRequest true "User request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users [post]
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var req request.CreateUsersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	user, err := uc.userService.CreateUser(req)
	if err != nil {
		ctx.Error(err)
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
// @Summary Get user by ID
// @Description Get user details by ID
// @Tags users_2get
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users/{id} [get]
func (uc *UserController) GetUser(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	user, err := uc.userService.GetUserByID(id)
	if err != nil {
		ctx.Error(err)
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
// @Description Update user details
// @Tags users_4update
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body request.UpdateUserRequest true "User request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users/{id} [put]
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var req request.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		return
	}
	updatedUser, err := uc.userService.UpdateUser(id, req)
	if err != nil {
		ctx.Error(err)
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
// @Tags users_3get
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users [get]
func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		ctx.Error(err)
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

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user by ID
// @Tags users_5delete
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users/{id} [delete]
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	err = uc.userService.DeleteUser(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "User deleted successfully",
	})
}
