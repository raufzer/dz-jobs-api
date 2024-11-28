package controllers

import (
	"net/http"

	"dz-jobs-api/config"
	"dz-jobs-api/data/request"
	"dz-jobs-api/data/response"
	"dz-jobs-api/internal/services"
	"dz-jobs-api/pkg/helpers"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	authService services.AuthenticationService
	config      *config.AppConfig
}

func NewAuthenticationController(service services.AuthenticationService, config *config.AppConfig) *AuthenticationController {
	return &AuthenticationController{
		authService: service,
		config:      config,
	}
}

func (ac *AuthenticationController) Login(ctx *gin.Context) {
	var loginRequest request.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		helpers.RespondWithError(ctx, "Invalid request format", http.StatusBadRequest)
		return
	}

	token, err := ac.authService.Login(loginRequest)
	if err != nil {
		helpers.RespondWithError(ctx, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	helpers.SetAuthCookie(ctx, token, ac.config.TokenMaxAge, "localhost")
	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}
	helpers.RespondWithSuccess(ctx, "Successfully logged in!", resp)
}

func (ac *AuthenticationController) Register(ctx *gin.Context) {
	var createUserRequest request.CreateUsersRequest
	if err := ctx.ShouldBindJSON(&createUserRequest); err != nil {
		helpers.RespondWithError(ctx, "Invalid request format", http.StatusBadRequest)
		return
	}

	if err := ac.authService.Register(createUserRequest); err != nil {
		helpers.RespondWithError(ctx, "Registration failed", http.StatusInternalServerError)
		return
	}

	helpers.RespondWithSuccess(ctx, "Successfully created user!", nil)
}
