package redis

import (
	"context"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces"
	"dz-jobs-api/pkg/utils"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisRepository struct {
	redisClient *redis.Client
}

func NewRedisRepository(redisClient *redis.Client) repositoryInterfaces.RedisRepository {
	return &RedisRepository{
		redisClient: redisClient,
	}
}

func (r *RedisRepository) StoreOTP(ctx context.Context, email, otp string, expiry time.Duration) error {
	key := fmt.Sprintf("otp:%s", email)
	if err := r.redisClient.Set(ctx, key, otp, expiry).Err(); err != nil {
		return fmt.Errorf("redis: failed to store OTP for email %s: %w", email, err)
	}
	return nil
}

func (r *RedisRepository) GetOTP(ctx context.Context, email string) (string, error) {
	key := fmt.Sprintf("otp:%s", email)
	result, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", redis.Nil
		}
		return "", fmt.Errorf("redis: failed to get OTP for email %s: %w", email, err)
	}
	return result, nil
}

func (r *RedisRepository) InvalidateOTP(ctx context.Context, email string) error {
	key := fmt.Sprintf("otp:%s", email)
	if err := r.redisClient.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("redis: failed to delete OTP for email %s: %w", email, err)
	}
	return nil
}

func (r *RedisRepository) StoreResetToken(ctx context.Context, email, token string, expiry time.Duration) error {
	key := fmt.Sprintf("reset_token:%s", email)
	if err := r.redisClient.Set(ctx, key, token, expiry).Err(); err != nil {
		return fmt.Errorf("redis: failed to store reset token for email %s: %w", email, err)
	}
	return nil
}

func (r *RedisRepository) GetResetToken(ctx context.Context, email string) (string, error) {
	key := fmt.Sprintf("reset_token:%s", email)
	result, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", redis.Nil
		}
		return "", fmt.Errorf("redis: failed to get reset token for email %s: %w", email, err)
	}
	return result, nil
}

func (r *RedisRepository) InvalidateResetToken(ctx context.Context, email string) error {
	key := fmt.Sprintf("reset_token:%s", email)
	if err := r.redisClient.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("redis: failed to delete reset token for email %s: %w", email, err)
	}
	return nil
}

func (r *RedisRepository) StoreRefreshToken(ctx context.Context, UserID, refreshToken string, expiry time.Duration) error {
	key := fmt.Sprintf("refresh_token:%s", UserID)
	if err := r.redisClient.Set(ctx, key, refreshToken, expiry).Err(); err != nil {
		return fmt.Errorf("redis: failed to store refresh token for UserID %s: %w", UserID, err)
	}
	return nil
}

func (r *RedisRepository) GetRefreshToken(ctx context.Context, UserID string) (string, error) {
	key := fmt.Sprintf("refresh_token:%s", UserID)
	result, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", redis.Nil
		}
		return "", fmt.Errorf("redis: failed to get refresh token for UserID %s: %w", UserID, err)
	}
	return result, nil
}

func (r *RedisRepository) InvalidateRefreshToken(ctx context.Context, UserID string) error {
	key := fmt.Sprintf("refresh_token:%s", UserID)
	if err := r.redisClient.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("redis: failed to delete refresh token for user_id %s: %w", UserID, err)
	}
	return nil
}

func (r *RedisRepository) StoreAssetCache(ctx context.Context, assetID string, assetType string, data *utils.AssetCache, expiry time.Duration) error {
	key := fmt.Sprintf("asset:%s:%s", assetType, assetID)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("redis: failed to marshal asset data: %w", err)
	}

	if err := r.redisClient.Set(ctx, key, jsonData, expiry).Err(); err != nil {
		return fmt.Errorf("redis: failed to store asset cache for ID %s: %w", assetID, err)
	}
	return nil
}

func (r *RedisRepository) GetAssetCache(ctx context.Context, assetID string, assetType string) (*utils.AssetCache, error) {
	key := fmt.Sprintf("asset:%s:%s", assetType, assetID)

	result, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("redis: failed to get asset cache for ID %s: %w", assetID, err)
	}

	var assetCache utils.AssetCache
	if err := json.Unmarshal([]byte(result), &assetCache); err != nil {
		return nil, fmt.Errorf("redis: failed to unmarshal asset data: %w", err)
	}

	return &assetCache, nil
}

func (r *RedisRepository) InvalidateAssetCache(ctx context.Context, assetID string, assetType string) error {
	key := fmt.Sprintf("asset:%s:%s", assetType, assetID)

	if err := r.redisClient.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("redis: failed to invalidate asset cache for ID %s: %w", assetID, err)
	}
	return nil
}
