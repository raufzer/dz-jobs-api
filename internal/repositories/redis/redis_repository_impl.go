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

func NewUserRepository(redisClient *redis.Client) repositoryInterfaces.RedisRepository {
	return &RedisRepository{
		redisClient: redisClient,
	}
}

func (r *RedisRepository) StoreOTP(email, otp string, expiry time.Duration) error {
	return r.redisClient.Set(context.Background(), email+"_otp", otp, expiry).Err()
}

func (r *RedisRepository) GetOTP(email string) (string, error) {
	return r.redisClient.Get(context.Background(), email+"_otp").Result()
}

func (r *RedisRepository) DeleteOTP(email string) error {
	return r.redisClient.Del(context.Background(), email+"_otp").Err()
}

func (r *RedisRepository) StoreResetToken(email, token string, expiry time.Duration) error {
	return r.redisClient.Set(context.Background(), email+"_reset_token", token, expiry).Err()
}

func (r *RedisRepository) GetResetToken(email string) (string, error) {
	return r.redisClient.Get(context.Background(), email+"_reset_token").Result()
}

func (r *RedisRepository) DeleteResetToken(email string) error {
	return r.redisClient.Del(context.Background(), email+"_reset_token").Err()
}
