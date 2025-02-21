package interfaces

import (
	"context"
	"dz-jobs-api/pkg/utils"
	"time"
)

type RedisRepository interface {
	StoreOTP(ctx context.Context, email, otp string, expiry time.Duration) error
	GetOTP(ctx context.Context, email string) (string, error)
	InvalidateOTP(ctx context.Context, email string) error
	StoreResetToken(ctx context.Context, email, token string, expiry time.Duration) error
	GetResetToken(ctx context.Context, email string) (string, error)
	InvalidateResetToken(ctx context.Context, email string) error
	StoreRefreshToken(ctx context.Context, UserID, refreshToken string, expiry time.Duration) error
	GetRefreshToken(ctx context.Context, UserID string) (string, error)
	InvalidateRefreshToken(ctx context.Context, UserID string) error
	StoreAssetCache(ctx context.Context, assetID string, assetType string, data *utils.AssetCache, expiry time.Duration) error
	GetAssetCache(ctx context.Context, assetID string, assetType string) (*utils.AssetCache, error)
	InvalidateAssetCache(ctx context.Context, assetID string, assetType string) error
}
