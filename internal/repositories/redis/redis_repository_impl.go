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

func (r *RedisRepository) StoreOTP(email, otp string, expiry time.Duration) error {
	key := fmt.Sprintf("otp:%s:", email)
	if err := r.redisClient.Set(context.Background(), key, otp, expiry).Err(); err != nil {
		return fmt.Errorf("redis: failed to store OTP for email %s: %w", email, err)
	}
	return nil
}

func (r *RedisRepository) GetOTP(email string) (string, error) {
	key := fmt.Sprintf("otp:%s:", email)
	result, err := r.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", redis.Nil
		}
		return "", fmt.Errorf("redis: failed to get OTP for email %s: %w", email, err)
	}
	return result, nil
}

func (r *RedisRepository) InvalidateOTP(email string) error {
	key := fmt.Sprintf("otp:%s:", email)
	if err := r.redisClient.Del(context.Background(), key).Err(); err != nil {
		return fmt.Errorf("redis: failed to delete OTP for email %s: %w", email, err)
	}
	return nil
}

func (r *RedisRepository) StoreResetToken(email, token string, expiry time.Duration) error {
	key := fmt.Sprintf("reset_token:%s:", email)
	if err := r.redisClient.Set(context.Background(), key, token, expiry).Err(); err != nil {
		return fmt.Errorf("redis: failed to store reset token for email %s: %w", email, err)
	}
	return nil
}

func (r *RedisRepository) GetResetToken(email string) (string, error) {
	key := fmt.Sprintf("reset_token:%s:", email)
	result, err := r.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", redis.Nil
		}
		return "", fmt.Errorf("redis: failed to get reset token for email %s: %w", email, err)
	}
	return result, nil
}

func (r *RedisRepository) InvalidateResetToken(email string) error {
	key := fmt.Sprintf("reset_token:%s:", email)
	if err := r.redisClient.Del(context.Background(), key).Err(); err != nil {
		return fmt.Errorf("redis: failed to delete reset token for email %s: %w", email, err)
	}
	return nil
}

func (r *RedisRepository) StoreRefreshToken(user_id, refreshToken string, expiry time.Duration) error {
	key := fmt.Sprintf("refresh_token:%s:", user_id)
	if err := r.redisClient.Set(context.Background(), key, refreshToken, expiry).Err(); err != nil {
		return fmt.Errorf("redis: failed to store refresh token for user_id %s: %w", user_id, err)
	}
	return nil
}

func (r *RedisRepository) GetRefreshToken(user_id string) (string, error) {
	key := fmt.Sprintf("refresh_token:%s:", user_id)
	result, err := r.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", redis.Nil
		}
		return "", fmt.Errorf("redis: failed to get refresh token for user_id %s: %w", user_id, err)
	}
	return result, nil
}

func (r *RedisRepository) InvalidateRefreshToken(user_id string) error {
	key := fmt.Sprintf("refresh_token:%s:", user_id)
	if err := r.redisClient.Del(context.Background(), key).Err(); err != nil {
		return fmt.Errorf("redis: failed to delete refresh token for user_id %s: %w", user_id, err)
	}
	return nil
}

func (r *RedisRepository) StoreAssetCache(assetID string, assetType string, data *utils.AssetCache, expiry time.Duration) error {
	key := fmt.Sprintf("asset:%s:%s", assetType, assetID)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("redis: failed to marshal asset data: %w", err)
	}

	if err := r.redisClient.Set(context.Background(), key, jsonData, expiry).Err(); err != nil {
		return fmt.Errorf("redis: failed to store asset cache for ID %s: %w", assetID, err)
	}
	return nil
}

func (r *RedisRepository) GetAssetCache(assetID string, assetType string) (*utils.AssetCache, error) {
	key := fmt.Sprintf("asset:%s:%s", assetType, assetID)

	result, err := r.redisClient.Get(context.Background(), key).Result()
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

func (r *RedisRepository) InvalidateAssetCache(assetID string, assetType string) error {
	key := fmt.Sprintf("asset:%s:%s", assetType, assetID)

	if err := r.redisClient.Del(context.Background(), key).Err(); err != nil {
		return fmt.Errorf("redis: failed to invalidate asset cache for ID %s: %w", assetID, err)
	}
	return nil
}
