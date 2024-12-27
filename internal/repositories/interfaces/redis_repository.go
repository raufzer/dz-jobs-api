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
	StoreRefreshToken(user_id, refreshToken string, expiry time.Duration) error
	GetRefreshToken(user_id string) (string, error)
	InvalidateRefreshToken(user_id string) error
	StoreAssetCache(assetID string, assetType string, data *utils.AssetCache, expiry time.Duration) error
	GetAssetCache(assetID string, assetType string) (*utils.AssetCache, error)
	InvalidateAssetCache(assetID string, assetType string) error
}
