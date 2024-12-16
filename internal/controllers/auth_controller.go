package controllers

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/dto/response"
	"dz-jobs-api/internal/helpers"
	serviceInterfaces "dz-jobs-api/internal/services/interfaces"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	accessToken, refreshToken, err := ac.authService.Login(req)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	isProduction := ac.config.ServerPort != "9090"
	helpers.SetAuthCookie(ctx, "access_token", accessToken, ac.config.AccessTokenMaxAge, ac.config.Domain, isProduction)
	helpers.SetAuthCookie(ctx, "refresh_token", refreshToken, ac.config.RefreshTokenMaxAge, ac.config.Domain, isProduction)
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully logged in!",
	})
}
func (ac *AuthController) RefreshToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	fmt.Println("Refresh token:", refreshToken)
	userID, userRole, err := ac.authService.ValidateToken(refreshToken)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	accessToken, err := ac.authService.RefreshAccessToken(userID, userRole, refreshToken)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	isProduction := ac.config.ServerPort != "9090"
	helpers.SetAuthCookie(ctx, "access_token", accessToken, ac.config.AccessTokenMaxAge, ac.config.Domain, isProduction)
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Access token refreshed successfully!",
	})
}

func (ac *AuthController) Logout(ctx *gin.Context) {
	isProduction := ac.config.ServerPort != "9090"
	helpers.SetAuthCookie(ctx, "access_token", "", -1, ac.config.Domain, isProduction)
	helpers.SetAuthCookie(ctx, "refresh_token", "", -1, ac.config.Domain, isProduction)
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully logged out!",
	})
}
func (ac *AuthController) Register(ctx *gin.Context) {
	var req request.CreateUsersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	if err := ac.authService.Register(req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "User successfully created!",
	})
}
func (ac *AuthController) SendResetOTP(ctx *gin.Context) {
	var req request.SendOTPRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	err := ac.authService.SendOTP(req.Email)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "OTP sent successfully!",
	})
}
func (ac *AuthController) VerifyOTP(ctx *gin.Context) {
	var req request.VerifyOTPRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	resetToken, err := ac.authService.VerifyOTP(req.Email, req.OTP)
	if err != nil {
		ctx.Error(err)
		return
	}
	isProduction := ac.config.ServerPort != "9090"
	helpers.SetAuthCookie(ctx, "reset_token", resetToken, ac.config.ResetPasswordTokenMaxAge, ac.config.Domain, isProduction)
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "OTP vefiy successfully!",
	})
}
func (ac *AuthController) ResetPassword(ctx *gin.Context) {
	token, err := ctx.Cookie("reset_token")
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	var req request.ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	err = ac.authService.ResetPassword(req.Email, token, req.NewPassword)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Password reset successfully!",
	})
}
