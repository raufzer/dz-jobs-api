package services

import (
	"context"
	"database/sql"
	"dz-jobs-api/config"
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/pkg/utils"
	"dz-jobs-api/internal/integrations"
	"dz-jobs-api/internal/models"
	"dz-jobs-api/internal/repositories/interfaces"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

type AuthService struct {
	userRepository  interfaces.UserRepository
	redisRepository interfaces.RedisRepository
	config          *config.AppConfig
}

func NewAuthService(userRepo interfaces.UserRepository, redisRepo interfaces.RedisRepository, config *config.AppConfig) *AuthService {
	return &AuthService{
		userRepository:  userRepo,
		redisRepository: redisRepo,
		config:          config,
	}
}
func (as *AuthService) Register(req request.CreateUsersRequest) (*models.User, error) {
	existingUser, _ := as.userRepository.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "User already exists")
	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to hash password")
	}
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.Role,
	}
	if err := as.userRepository.Create(user); err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to create user")
	}
	return user, nil
}
func (as *AuthService) Login(req request.LoginRequest) (*models.User, string, string, error) {
	user, err := as.userRepository.GetByEmail(req.Email)
	if err != nil || user == nil {
		return nil, "", "", utils.NewCustomError(http.StatusUnauthorized, "Invalid email")
	}

	verifyErr := utils.VerifyPassword(user.Password, req.Password)
	if verifyErr != nil {
		return nil, "", "", utils.NewCustomError(http.StatusUnauthorized, "Invalid password")
	}

	accessToken, err := utils.GenerateToken(user.ID.String(), as.config.AccessTokenMaxAge, "access", user.Role, as.config.AccessTokenSecret)
	if err != nil {
		return nil, "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to generate access token")
	}
	refreshToken, err := utils.GenerateToken(user.ID.String(), as.config.RefreshTokenMaxAge, "refresh", "", as.config.RefreshTokenSecret)
	if err != nil {
		return nil, "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to generate refresh token")
	}
	refreshTokenTTL := as.config.RefreshTokenMaxAge
	err = as.redisRepository.StoreRefreshToken(user.ID.String(), refreshToken, refreshTokenTTL)
	if err != nil {
		return nil, "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to store refresh token")
	}

	return user, accessToken, refreshToken, nil
}

func (as *AuthService) RefreshAccessToken(userID, userRole, refreshToken string) (string, error) {
	storedToken, err := as.redisRepository.GetRefreshToken(userID)
	if err != nil {
		if err == redis.Nil {
			return "", utils.NewCustomError(http.StatusUnauthorized, "Refresh token expired or not found")
		}
		return "", utils.NewCustomError(http.StatusInternalServerError, "Failed to retrieve refresh token")
	}
	if storedToken != refreshToken {
		return "", utils.NewCustomError(http.StatusUnauthorized, "Invalid Token")
	}

	accessToken, err := utils.GenerateToken(userID, as.config.AccessTokenMaxAge, "access", userRole, as.config.AccessTokenSecret)
	if err != nil {
		return "", utils.NewCustomError(http.StatusInternalServerError, "Failed to generate access token")
	}

	return accessToken, nil
}

func (as *AuthService) SendOTP(email string) error {
	otp := utils.GenerateSecureOTP(6)
	err := as.redisRepository.StoreOTP(email, otp, 5*time.Minute)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to store OTP")
	}

	return integrations.SendOTPEmail(email, otp, as.config.SendGridAPIKey)
}

func (as *AuthService) VerifyOTP(email, otp string) (string, error) {
	storedOTP, err := as.redisRepository.GetOTP(email)
	if err != nil {
		if err == redis.Nil {
			return "", utils.NewCustomError(http.StatusUnauthorized, "OTP expired or not found")
		}
		return "", utils.NewCustomError(http.StatusInternalServerError, "Failed to retrieve OTP")
	}
	if storedOTP != otp {
		return "", utils.NewCustomError(http.StatusUnauthorized, "Invalid OTP")
	}
	resetToken, err := utils.GenerateToken(email, as.config.AccessTokenMaxAge, "reset_password", "", as.config.ResetPasswordTokenSecret)
	if err != nil {
		return "", utils.NewCustomError(http.StatusInternalServerError, "Failed to generate reset password token")
	}
	err = as.redisRepository.StoreResetToken(email, resetToken, as.config.ResetPasswordTokenMaxAge)
	if err != nil {
		return "", utils.NewCustomError(http.StatusInternalServerError, "Failed to store reset token")
	}

	_ = as.redisRepository.DeleteOTP(email)

	return resetToken, nil
}

func (as *AuthService) ResetPassword(email, resetToken, newPassword string) error {
	storedToken, err := as.redisRepository.GetResetToken(email)
	if err != nil {
		if err == redis.Nil {
			return utils.NewCustomError(http.StatusUnauthorized, "Reset token expired or not found")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to retrieve reset token")
	}
	if storedToken != resetToken {
		return utils.NewCustomError(http.StatusUnauthorized, "Invalid reset token")
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to hash new password")
	}

	err = as.userRepository.UpdatePassword(email, string(hashedPassword))
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to update password")
	}

	_ = as.redisRepository.DeleteResetToken(email)

	return nil
}

func (as *AuthService) ValidateToken(token string) (string, string, error) {

	claims, err := utils.ValidateToken(token, as.config.RefreshTokenSecret, "refresh")
	if err != nil {
		return "", "", utils.NewCustomError(http.StatusUnauthorized, "Invalid or expired token")
	}
	return claims.UserID, claims.Role, nil
}

func (as *AuthService) GoogleConnect(code string, role string) (*models.User, string, string, string, error) {

	oauthConfig := integrations.InitializeGoogleOAuthConfig(as.config.GoogleClientID, as.config.GoogleClientSecret, as.config.GoogleRedirectURL)

	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, "", "", "", utils.NewCustomError(http.StatusBadRequest, "Failed to exchange authorization code for token")
	}

	userInfo, err := integrations.FetchGoogleUserInfo(oauthConfig, token)
	if err != nil {
		return nil, "", "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch user information from Google")
	}

	existingUser, err := as.userRepository.GetByEmail(userInfo.Email)
	if err != nil {

		if err == sql.ErrNoRows {
			hashedPassword, err := utils.HashPassword(utils.GenerateRandomPassword())
			if err != nil {
				return nil, "", "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to hash password")
			}
			newUser := &models.User{
				Name:     userInfo.Name,
				Email:    userInfo.Email,
				Role:     role,
				Password: hashedPassword,
			}

			if err := as.userRepository.Create(newUser); err != nil {
				return nil, "", "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to create new user")
			}
			return newUser, "", "", "register", nil
		}

		return nil, "", "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to check user existence")
	}
	accessToken, err := utils.GenerateToken(userInfo.ID, as.config.AccessTokenMaxAge, "access", role, as.config.AccessTokenSecret)
	if err != nil {
		return nil, "", "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to generate access token")
	}
	refreshToken, err := utils.GenerateToken(userInfo.ID, as.config.RefreshTokenMaxAge, "refresh", "", as.config.RefreshTokenSecret)
	if err != nil {
		return nil, "", "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to generate refresh token")
	}
	refreshTokenTTL := as.config.RefreshTokenMaxAge
	err = as.redisRepository.StoreRefreshToken(userInfo.ID, refreshToken, refreshTokenTTL)
	if err != nil {
		return nil, "", "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to store refresh token")
	}
	return existingUser, accessToken, refreshToken, "login", nil
}
