package interfaces

import (
	"dz-jobs-api/pkg/utils"
	"time"
)

type RedisRepository interface {
	StoreOTP(email, otp string, expiry time.Duration) error
	GetOTP(email string) (string, error)
	InvalidateOTP(email string) error
	StoreResetToken(email, token string, expiry time.Duration) error
	GetResetToken(email string) (string, error)
	InvalidateResetToken(email string) error
	StoreRefreshToken(UserID, refreshToken string, expiry time.Duration) error
	GetRefreshToken(UserID string) (string, error)
	InvalidateRefreshToken(UserID string) error
	StoreAssetCache(assetID string, assetType string, data *utils.AssetCache, expiry time.Duration) error
	GetAssetCache(assetID string, assetType string) (*utils.AssetCache, error)
	InvalidateAssetCache(assetID string, assetType string) error
}
