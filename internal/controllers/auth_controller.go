package controllers

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/dto/response"
	"dz-jobs-api/internal/integrations"
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
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "Login request"
// @Success 200 {object} response.Response{Data=response.UserResponse} "Successfully logged in!"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	user, accessToken, refreshToken, err := c.authService.Login(req)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	isProduction := c.config.ServerPort != "9090"
	utils.SetAuthCookie(ctx, "access_token", accessToken, c.config.AccessTokenMaxAge, c.config.BackEndDomain, isProduction)
	utils.SetAuthCookie(ctx, "refresh_token", refreshToken, c.config.RefreshTokenMaxAge, c.config.BackEndDomain, isProduction)
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
// @Tags Auth
// @Produce json
// @Success 200 {object} response.Response "Access token refreshed successfully!"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /auth/refresh-token [post]
func (c *AuthController) RefreshToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	userID, userRole, err := c.authService.ValidateToken(refreshToken)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	accessToken, err := c.authService.RefreshAccessToken(userID, userRole, refreshToken)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	isProduction := c.config.ServerPort != "9090"
	utils.SetAuthCookie(ctx, "access_token", accessToken, c.config.AccessTokenMaxAge, c.config.BackEndDomain, isProduction)
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Access token refreshed successfully!",
	})
}

// Logout godoc
// @Summary Logout user
// @Description Logout user and clear cookies
// @Tags Auth
// @Produce json
// @Success 200 {object} response.Response "Successfully logged out!"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /auth/logout [post]
func (c *AuthController) Logout(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	userID, _, _ := c.authService.ValidateToken(refreshToken)
	c.authService.Logout(userID, refreshToken)
	isProduction := c.config.ServerPort != "9090"
	utils.SetAuthCookie(ctx, "access_token", "", -1, c.config.BackEndDomain, isProduction)
	utils.SetAuthCookie(ctx, "refresh_token", "", -1, c.config.BackEndDomain, isProduction)

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully logged out!",
	})
}

// Register godoc
// @Summary Register user
// @Description Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.CreateUsersRequest true "Register request"
// @Success 201 {object} response.Response{Data=response.UserResponse} "User created successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 409 {object} response.Response "User already exists"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var req request.CreateUsersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	user, err := c.authService.Register(req)
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
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.SendOTPRequest true "Send OTP request"
// @Success 200 {object} response.Response "OTP sent successfully!"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /auth/send-reset-otp [post]
func (c *AuthController) SendResetOTP(ctx *gin.Context) {
	var req request.SendOTPRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	err := c.authService.SendOTP(req.Email)
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
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.VerifyOTPRequest true "Verify OTP request"
// @Success 200 {object} response.Response "OTP verify successfully!"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /auth/verify-otp [post]
func (c *AuthController) VerifyOTP(ctx *gin.Context) {
	var req request.VerifyOTPRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	resetToken, err := c.authService.VerifyOTP(req.Email, req.OTP)
	if err != nil {
		ctx.Error(err)
		return
	}
	isProduction := c.config.ServerPort != "9090"
	utils.SetAuthCookie(ctx, "reset_token", resetToken, c.config.ResetPasswordTokenMaxAge, c.config.BackEndDomain, isProduction)
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "OTP verify successfully!",
	})
}

// ResetPassword godoc
// @Summary Reset password
// @Description Reset user's password using reset token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.ResetPasswordRequest true "Reset password request"
// @Success 200 {object} response.Response "Password reset successfully!"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /auth/reset-password [post]
func (c *AuthController) ResetPassword(ctx *gin.Context) {
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
	err = c.authService.ResetPassword(req.Email, token, req.NewPassword)
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
// @Description Connect with Google OAuth (Register or Login)
// @Tags Auth
// @Param role query string true "User role (admin, candidate, recruiter)"
// @Produce json
// @Success 200 {object} response.Response{Data=response.UserResponse} "Successfully logged in!"
// @Success 201 {object} response.Response{Data=response.UserResponse} ""User created successfully""
// @Failure 400 {object} response.Response "Role is missing or expired"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /auth/google/connect [get]
func (c *AuthController) GoogleConnect(ctx *gin.Context) {
	role := ctx.Query("role")
	ctx.SetCookie("role", role, 3600, "/", c.config.BackEndDomain, false, true)
	oauthConfig := integrations.InitializeGoogleOAuthConfig(c.config.GoogleClientID, c.config.GoogleClientSecret, c.config.GoogleRedirectURL)

	authURL := oauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)

	ctx.Redirect(http.StatusFound, authURL)
}

func (c *AuthController) GoogleCallbackConnect(ctx *gin.Context) {
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

	user, accessToken, refreshToken, connect, err := c.authService.GoogleConnect(code, role)
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
		isProduction := c.config.ServerPort != "9090"
		utils.SetAuthCookie(ctx, "access_token", accessToken, c.config.AccessTokenMaxAge, c.config.BackEndDomain, isProduction)
		utils.SetAuthCookie(ctx, "refresh_token", refreshToken, c.config.RefreshTokenMaxAge, c.config.BackEndDomain, isProduction)
		ctx.JSON(http.StatusOK, response.Response{
			Code:    http.StatusOK,
			Status:  "OK",
			Message: "Successfully logged in!",
			Data:    response.ToUserResponse(user),
		})

	}
}
