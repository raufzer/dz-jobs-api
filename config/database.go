package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	DB *sql.DB
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
