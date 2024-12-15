package interfaces

import (
	"time"
)

type RedisRepository interface {
	StoreOTP(email, otp string, expiry time.Duration) error
	GetOTP(email string) (string, error)
	DeleteOTP(email string) error
	StoreResetToken(email, token string, expiry time.Duration) error
	GetResetToken(email string) (string, error)
	DeleteResetToken(email string) error
	StoreRefreshToken(email, refreshToken string, expiry time.Duration) error
	GetRefreshToken(email string) (string, error)
	DeleteRefreshToken(email string) error
	StoreAccessToken(email, accessToken string, expiry time.Duration) error
	GetAccessToken(email string) (string, error)
	DeleteAccessToken(email string) error
}
