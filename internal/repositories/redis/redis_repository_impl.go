package redis

import (
	"context"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces"
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
	key := email + ":otp"
	if err := r.redisClient.Set(context.Background(), key, otp, expiry).Err(); err != nil {
		return fmt.Errorf("redis: failed to store OTP for email %s: %w", email, err)
	}
	return nil
}

func (r *RedisRepository) GetOTP(email string) (string, error) {
	key := email + ":otp"
	result, err := r.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", redis.Nil // Key not found
		}
		return "", fmt.Errorf("redis: failed to get OTP for email %s: %w", email, err)
	}
	return result, nil
}

func (r *RedisRepository) DeleteOTP(email string) error {
	key := email + ":otp"
	if err := r.redisClient.Del(context.Background(), key).Err(); err != nil {
		return fmt.Errorf("redis: failed to delete OTP for email %s: %w", email, err)
	}
	return nil
}

func (r *RedisRepository) StoreResetToken(email, token string, expiry time.Duration) error {
	key := email + ":reset_token"
	if err := r.redisClient.Set(context.Background(), key, token, expiry).Err(); err != nil {
		return fmt.Errorf("redis: failed to store reset token for email %s: %w", email, err)
	}
	return nil
}

func (r *RedisRepository) GetResetToken(email string) (string, error) {
	key := email + ":reset_token"
	result, err := r.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", redis.Nil
		}
		return "", fmt.Errorf("redis: failed to get reset token for email %s: %w", email, err)
	}
	return result, nil
}

func (r *RedisRepository) DeleteResetToken(email string) error {
	key := email + ":reset_token"
	if err := r.redisClient.Del(context.Background(), key).Err(); err != nil {
		return fmt.Errorf("redis: failed to delete reset token for email %s: %w", email, err)
	}
	return nil
}

func (r *RedisRepository) StoreRefreshToken(user_id, refreshToken string, expiry time.Duration) error {
	key := user_id + ":refresh_token"
	if err := r.redisClient.Set(context.Background(), key, refreshToken, expiry).Err(); err != nil {
		return fmt.Errorf("redis: failed to store refresh token for user_id %s: %w", user_id, err)
	}
	return nil
}

func (r *RedisRepository) GetRefreshToken(user_id string) (string, error) {
	key := user_id + ":refresh_token"
	result, err := r.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", redis.Nil
		}
		return "", fmt.Errorf("redis: failed to get refresh token for user_id %s: %w", user_id, err)
	}
	return result, nil
}

func (r *RedisRepository) DeleteRefreshToken(user_id string) error {
	key := user_id + ":refresh_token"
	if err := r.redisClient.Del(context.Background(), key).Err(); err != nil {
		return fmt.Errorf("redis: failed to delete refresh token for user_id %s: %w", user_id, err)
	}
	return nil
}
