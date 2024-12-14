package config

import (
	"context"
	"database/sql"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	DB *sql.DB
}

type RedisConfig struct {
	Client *redis.Client
}

func ConnectDatabase(config *AppConfig) *DatabaseConfig {
	db, err := sql.Open("postgres", config.DatabaseURI)
	if err != nil {
		log.Fatalf("Error while opening connection to PostgreSQL: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error while connecting to PostgreSQL: %v", err)
	}

	log.Println("Connected to PostgreSQL successfully")

	return &DatabaseConfig{
		DB: db,
	}
}

func ConnectRedis(config *AppConfig) *RedisConfig {

	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisURI,
		Password: config.RedisPassword,
		DB:       0,
	})

	ctx := context.Background()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Error while connecting to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully")

	return &RedisConfig{
		Client: client,
	}
}
