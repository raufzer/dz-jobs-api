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
	StoreRefreshToken(userid, refreshToken string, expiry time.Duration) error
	GetRefreshToken(userid string) (string, error)
	DeleteRefreshToken(userid string) error

}
