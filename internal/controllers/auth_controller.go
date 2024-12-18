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

// @Summary Login to the system
// @Description Logs the user in and returns access and refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param loginRequest body request.LoginRequest true "Login request"
// @Success 200 {object} response.Response "Login successful"
// @Failure 400 {object} response.Response "Invalid input"
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

// @Summary Refresh the access token
// @Description Refreshes the user's access token using the refresh token stored in cookies
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "Access token refreshed"
// @Failure 400 {object} response.Response "Error
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

// @Summary Logout the user
// @Description Logs the user out and removes authentication cookies
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "Logged out successfully"
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

// @Summary Register a new user
// @Description Creates a new user in the system
// @Tags Auth
// @Accept json
// @Produce json
// @Param createUserRequest body request.CreateUsersRequest true "User registration request"
// @Success 201 {object} response.Response "User created successfully"
// @Failure 400 {object} response.Response "Error in user creation"
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

// @Summary Send OTP for password reset
// @Description Sends an OTP to the user's email for password reset
// @Tags Auth
// @Accept json
// @Produce json
// @Param sendOTPRequest body request.SendOTPRequest true "Send OTP request"
// @Success 200 {object} response.Response "OTP sent successfully"
// @Failure 400 {object} response.Response "Error in sending OTP"
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

// @Summary Verify OTP for password reset
// @Description Verifies the OTP sent to the user's email for password reset
// @Tags Auth
// @Accept json
// @Produce json
// @Param verifyOTPRequest body request.VerifyOTPRequest true "Verify OTP request"
// @Success 200 {object} response.Response "OTP verified successfully"
// @Failure 400 {object} response.Response "Error in OTP verification"
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
		Message: "OTP vefiy successfully!",
	})
}

// @Summary Reset the user's password
// @Description Resets the user's password using the provided token and new password
// @Tags Auth
// @Accept json
// @Produce json
// @Param resetPasswordRequest body request.ResetPasswordRequest true "Reset password request"
// @Success 200 {object} response.Response "Password reset successfully"
// @Failure 400 {object} response.Response "Error in password reset"
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

// @Summary Connect with Google OAuth
// @Description Initiates the Google OAuth flow
// @Tags Auth
// @Accept json
// @Produce json
// @Success 302 {object} response.Response "Redirecting to Google OAuth"
// @Router /auth/google/connect [get]
func (ac *AuthController) GoogleConnect(ctx *gin.Context) {
	role := ctx.DefaultQuery("role", "candidate")
	ctx.SetCookie("role", role, 3600, "/", ac.config.Domain, false, true)
	oauthConfig := utils.InitializeGoogleOAuthConfig(ac.config.GoogleClientID, ac.config.GoogleClientSecret, ac.config.GoogleRedirectURL)

	authURL := oauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)

	ctx.Redirect(http.StatusFound, authURL)
}

// @Summary Handle Google OAuth callback
// @Description Handles the Google OAuth callback and logs the user in or registers the user
// @Tags Auth
// @Accept json
// @Produce json
// @Param code query string true "OAuth code"
// @Success 200 {object} response.Response "Successfully connected via Google"
// @Failure 400 {object} response.Response "Error in OAuth callback"
// @Router /auth/google/callback [get]
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
