package services

import (
    "context"
    "database/sql"
    "dz-jobs-api/config"
    "dz-jobs-api/internal/dto/request"
    "dz-jobs-api/internal/integrations"
    "dz-jobs-api/internal/models"
    "dz-jobs-api/internal/repositories/interfaces"
    "dz-jobs-api/pkg/utils"
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

func (s *AuthService) Register(ctx context.Context, req request.CreateUsersRequest) (*models.User, error) {
    existingUser, err := s.userRepository.GetUserByEmail(ctx, req.Email)
    if err != nil {
        if err == sql.ErrNoRows {
            existingUser = nil
        }
    }
    if existingUser != nil {
        return nil, utils.NewCustomError(http.StatusBadRequest, "User already exists")
    } else {
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
        if err := s.userRepository.CreateUser(ctx, user); err != nil {
            return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to create user")
        }
        return user, nil
    }
}

func (s *AuthService) Login(ctx context.Context, req request.LoginRequest) (*models.User, string, string, error) {
    user, err := s.userRepository.GetUserByEmail(ctx, req.Email)
    if err != nil || user == nil {
        if err == sql.ErrNoRows {
            return nil, "", "", utils.NewCustomError(http.StatusUnauthorized, "Invalid email or password")
        }
        return nil, "", "", utils.NewCustomError(http.StatusUnauthorized, "Invalid email")
    }

    verifyErr := utils.VerifyPassword(user.Password, req.Password)
    if verifyErr != nil {
        return nil, "", "", utils.NewCustomError(http.StatusUnauthorized, "Invalid password")
    }

    accessToken, err := utils.GenerateToken(user.ID.String(), s.config.AccessTokenMaxAge, "access", user.Role, s.config.AccessTokenSecret)
    if err != nil {
        return nil, "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to generate access token")
    }
    refreshToken, err := utils.GenerateToken(user.ID.String(), s.config.RefreshTokenMaxAge, "refresh", "", s.config.RefreshTokenSecret)
    if err != nil {
        return nil, "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to generate refresh token")
    }
    refreshTokenTTL := s.config.RefreshTokenMaxAge
    err = s.redisRepository.StoreRefreshToken(ctx, user.ID.String(), refreshToken, refreshTokenTTL)
    if err != nil {
        return nil, "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to store refresh token")
    }

    return user, accessToken, refreshToken, nil
}

func (s *AuthService) RefreshAccessToken(ctx context.Context, userID, userRole, refreshToken string) (string, error) {
    storedToken, err := s.redisRepository.GetRefreshToken(ctx, userID)
    if err != nil {
        if err == redis.Nil {
            return "", utils.NewCustomError(http.StatusUnauthorized, "Refresh token expired or not found")
        }
        return "", utils.NewCustomError(http.StatusInternalServerError, "Failed to retrieve refresh token")
    }
    if storedToken != refreshToken {
        return "", utils.NewCustomError(http.StatusUnauthorized, "Invalid Token")
    }

    accessToken, err := utils.GenerateToken(userID, s.config.AccessTokenMaxAge, "access", userRole, s.config.AccessTokenSecret)
    if err != nil {
        return "", utils.NewCustomError(http.StatusInternalServerError, "Failed to generate access token")
    }

    return accessToken, nil
}

func (s *AuthService) Logout(ctx context.Context, userID, refreshToken string) error {
    err := s.redisRepository.InvalidateRefreshToken(ctx, userID)
    if err != nil {
        return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete refresh token")
    }
    return nil
}

func (s *AuthService) SendOTP(ctx context.Context, email string) error {
    otp := utils.GenerateSecureOTP(6)
    err := s.redisRepository.StoreOTP(ctx, email, otp, 5*time.Minute)
    if err != nil {
        return utils.NewCustomError(http.StatusInternalServerError, "Failed to store OTP")
    }

    return integrations.SendOTPEmail(email, otp, s.config.SendGridAPIKey)
}

func (s *AuthService) VerifyOTP(ctx context.Context, email, otp string) (string, error) {
    storedOTP, err := s.redisRepository.GetOTP(ctx, email)
    if err != nil {
        if err == redis.Nil {
            return "", utils.NewCustomError(http.StatusUnauthorized, "OTP expired or not found")
        }
        return "", utils.NewCustomError(http.StatusInternalServerError, "Failed to retrieve OTP")
    }
    if storedOTP != otp {
        return "", utils.NewCustomError(http.StatusUnauthorized, "Invalid OTP")
    }
    resetToken, err := utils.GenerateToken(email, s.config.AccessTokenMaxAge, "reset_password", "", s.config.ResetPasswordTokenSecret)
    if err != nil {
        return "", utils.NewCustomError(http.StatusInternalServerError, "Failed to generate reset password token")
    }
    err = s.redisRepository.StoreResetToken(ctx, email, resetToken, s.config.ResetPasswordTokenMaxAge)
    if err != nil {
        return "", utils.NewCustomError(http.StatusInternalServerError, "Failed to store reset token")
    }

    if err := s.redisRepository.InvalidateOTP(ctx, email); err != nil {
        return "", utils.NewCustomError(http.StatusInternalServerError, "Failed to delete OTP")
    }

    return resetToken, nil
}

func (s *AuthService) ResetPassword(ctx context.Context, email, resetToken, newPassword string) error {
    storedToken, err := s.redisRepository.GetResetToken(ctx, email)
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

    err = s.userRepository.UpdateUserPassword(ctx, email, string(hashedPassword))
    if err != nil {
        return utils.NewCustomError(http.StatusInternalServerError, "Failed to update password")
    }

    _ = s.redisRepository.InvalidateResetToken(ctx, email)

    return nil
}

func (s *AuthService) ValidateToken(ctx context.Context, token string) (string, string, error) {
    claims, err := utils.ValidateToken(token, s.config.RefreshTokenSecret, "refresh")
    if err != nil {
        return "", "", utils.NewCustomError(http.StatusUnauthorized, "Invalid or expired token")
    }
    return claims.ID, claims.Role, nil
}

func (s *AuthService) GoogleConnect(ctx context.Context, code string, role string) (*models.User, string, string, string, error) {
    oauthConfig := integrations.InitializeGoogleOAuthConfig(s.config.GoogleClientID, s.config.GoogleClientSecret, s.config.GoogleRedirectURL)

    token, err := oauthConfig.Exchange(ctx, code)
    if err != nil {
        return nil, "", "", "", utils.NewCustomError(http.StatusBadRequest, "Failed to exchange authorization code for token")
    }

    userInfo, err := integrations.FetchGoogleUserInfo(oauthConfig, token)
    if err != nil {
        return nil, "", "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch user information from Google")
    }

    existingUser, err := s.userRepository.GetUserByEmail(ctx, userInfo.Email)
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

            if err := s.userRepository.CreateUser(ctx, newUser); err != nil {
                return nil, "", "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to create new user")
            }
            return newUser, "", "", "register", nil
        }

        return nil, "", "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to check user existence")
    }
    accessToken, err := utils.GenerateToken(userInfo.ID, s.config.AccessTokenMaxAge, "access", role, s.config.AccessTokenSecret)
    if err != nil {
        return nil, "", "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to generate access token")
    }
    refreshToken, err := utils.GenerateToken(userInfo.ID, s.config.RefreshTokenMaxAge, "refresh", "", s.config.RefreshTokenSecret)
    if err != nil {
        return nil, "", "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to generate refresh token")
    }
    refreshTokenTTL := s.config.RefreshTokenMaxAge
    err = s.redisRepository.StoreRefreshToken(ctx, userInfo.ID, refreshToken, refreshTokenTTL)
    if err != nil {
        return nil, "", "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to store refresh token")
    }
    return existingUser, accessToken, refreshToken, "login", nil
}