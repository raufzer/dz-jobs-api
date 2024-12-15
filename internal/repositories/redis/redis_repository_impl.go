package redis

import (
	"context"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces"
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
	return r.redisClient.Set(context.Background(), key, otp, expiry).Err()
}

func (r *RedisRepository) GetOTP(email string) (string, error) {
	key := email + ":otp"
	return r.redisClient.Get(context.Background(), key).Result()
}

func (r *RedisRepository) DeleteOTP(email string) error {
	key := email + ":otp"
	return r.redisClient.Del(context.Background(), key).Err()
}

func (r *RedisRepository) StoreResetToken(email, token string, expiry time.Duration) error {
	key := email + ":reset_token"
	return r.redisClient.Set(context.Background(), key, token, expiry).Err()
}

func (r *RedisRepository) GetResetToken(email string) (string, error) {
	key := email + ":reset_token"
	return r.redisClient.Get(context.Background(), key).Result()
}

func (r *RedisRepository) DeleteResetToken(email string) error {
	key := email + ":reset_token"
	return r.redisClient.Del(context.Background(), key).Err()
}

func (r *RedisRepository) StoreRefreshToken(email, refreshToken string, expiry time.Duration) error {
	key := email + ":refresh_token"
	return r.redisClient.Set(context.Background(), key, refreshToken, expiry).Err()
}

func (r *RedisRepository) GetRefreshToken(email string) (string, error) {
	key := email + ":refresh_token"
	return r.redisClient.Get(context.Background(), key).Result()
}

func (r *RedisRepository) DeleteRefreshToken(email string) error {
	key := email + ":refresh_token"
	return r.redisClient.Del(context.Background(), key).Err()
}

func (r *RedisRepository) StoreAccessToken(email, accessToken string, expiry time.Duration) error {
	key := email + ":access_token"
	return r.redisClient.Set(context.Background(), key, accessToken, expiry).Err()
}

func (r *RedisRepository) GetAccessToken(email string) (string, error) {
	key := email + ":access_token"
	return r.redisClient.Get(context.Background(), key).Result()
}

func (r *RedisRepository) DeleteAccessToken(email string) error {
	key := email + ":access_token"
	return r.redisClient.Del(context.Background(), key).Err()
}
