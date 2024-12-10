package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"dz-jobs-api/config"
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/dto/response"
	"dz-jobs-api/internal/helpers"
	serviceInterfaces "dz-jobs-api/internal/services/interfaces"
)

type AuthController struct {
	authService serviceInterfaces.AuthService
	config      *config.AppConfig
}

func NewAuthController(service serviceInterfaces.AuthService, config *config.AppConfig) *AuthController {
	return &AuthController{
		authService: service,
		config:      config,
	}
}

func (ac *AuthController) Login(ctx *gin.Context) {
	var loginRequest request.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
        ctx.Error(err)
		return
	}

	token, err := ac.authService.Login(loginRequest)
	if err != nil {
		ctx.Error(err)
		return
	}

	isProduction := ac.config.ServerPort != "9090"
	helpers.SetAuthCookie(ctx, token, ac.config.TokenMaxAge, ac.config.Domain, isProduction)

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully logged in!",
	})
}

func (ac *AuthController) Register(ctx *gin.Context) {
	var createUserRequest request.CreateUsersRequest
	if err := ctx.ShouldBindJSON(&createUserRequest); err != nil {
        ctx.Error(err)
		return
	}

	if err := ac.authService.Register(createUserRequest); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "User successfully created!",
	})
}
