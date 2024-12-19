package controllers

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/dto/response"
	"dz-jobs-api/internal/helpers"
	serviceInterfaces "dz-jobs-api/internal/services/interfaces"
	"dz-jobs-api/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// AuthController handles authentication related operations
type AuthController struct {
	authService serviceInterfaces.AuthService
	config      *config.AppConfig
}

// NewAuthController creates a new AuthController
func NewAuthController(service serviceInterfaces.AuthService, config *config.AppConfig) *AuthController {
	return &AuthController{
		authService: service,
		config:      config,
	}
}

// Login godoc
// @Summary Login user
// @Description Login user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "Login request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /auth/login [post]
func (ac *AuthController) Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	user, accessToken, refreshToken, err := ac.authService.Login(req)
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
		Data:    response.ToUserResponse(user),
	})
}

// RefreshToken godoc
// @Summary Refresh access token
// @Description Refresh access token using refresh token
// @Tags auth
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /auth/refresh-token [post]
func (ac *AuthController) RefreshToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
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

// Logout godoc
// @Summary Logout user
// @Description Logout user and clear cookies
// @Tags auth
// @Produce json
// @Success 200 {object} response.Response
// @Router /auth/logout [post]
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

// Register godoc
// @Summary Register user
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body request.CreateUsersRequest true "Register request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /auth/register [post]
func (ac *AuthController) Register(ctx *gin.Context) {
	var req request.CreateUsersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	user, err := ac.authService.Register(req)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "User created successfully",
		Data:    response.ToUserResponse(user),
	})
}

// SendResetOTP godoc
// @Summary Send OTP for password reset
// @Description Send OTP to user's email for password reset
// @Tags auth
// @Accept json
// @Produce json
// @Param request body request.SendOTPRequest true "Send OTP request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /auth/send-reset-otp [post]
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

// VerifyOTP godoc
// @Summary Verify OTP
// @Description Verify OTP and generate reset token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body request.VerifyOTPRequest true "Verify OTP request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /auth/verify-otp [post]
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
		Message: "OTP verify successfully!",
	})
}

// ResetPassword godoc
// @Summary Reset password
// @Description Reset user's password using reset token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body request.ResetPasswordRequest true "Reset password request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /auth/reset-password [post]
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

// GoogleConnect godoc
// @Summary Google OAuth Connect
// @Description Connect with Google OAuth
// @Tags auth
// @Produce json
// @Success 302
// @Router /auth/google-connect [get]
func (ac *AuthController) GoogleConnect(ctx *gin.Context) {
	role := ctx.DefaultQuery("role", "candidate")
	ctx.SetCookie("role", role, 3600, "/", ac.config.Domain, false, true)
	oauthConfig := utils.InitializeGoogleOAuthConfig(ac.config.GoogleClientID, ac.config.GoogleClientSecret, ac.config.GoogleRedirectURL)

	authURL := oauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)

	ctx.Redirect(http.StatusFound, authURL)
}

func (ac *AuthController) GoogleCallbackConnect(ctx *gin.Context) {
	role, err := ctx.Cookie("role")
	if err != nil || role == "" {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Role is missing or expired",
		})
		return
	}

	code := ctx.DefaultQuery("code", "")

	if code == "" {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Code is required",
		})
		return
	}

	user, accessToken, refreshToken, connect, err := ac.authService.GoogleConnect(code, role)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	if connect == "register" {
		ctx.JSON(http.StatusOK, response.Response{
			Code:    http.StatusOK,
			Status:  "OK",
			Message: "User successfully created!",
			Data:    response.ToUserResponse(user),
		})
	} else if connect == "login" {
		isProduction := ac.config.ServerPort != "9090"
		helpers.SetAuthCookie(ctx, "access_token", accessToken, ac.config.AccessTokenMaxAge, ac.config.Domain, isProduction)
		helpers.SetAuthCookie(ctx, "refresh_token", refreshToken, ac.config.RefreshTokenMaxAge, ac.config.Domain, isProduction)
		ctx.JSON(http.StatusOK, response.Response{
			Code:    http.StatusOK,
			Status:  "OK",
			Message: "Successfully logged in!",
			Data:    response.ToUserResponse(user),
		})

	}
}
