package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type DatabaseConfig struct {
	DB *sql.DB
}

// ConnectDatabase initializes and connects to the PostgreSQL database using raw SQL.
func ConnectDatabase(config *AppConfig) *DatabaseConfig {
	// Open a connection to PostgreSQL
	db, err := sql.Open("postgres", config.DatabaseURI)
	if err != nil {
		log.Fatalf("Error while opening connection to PostgreSQL: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Error while connecting to PostgreSQL: %v", err)
	}

	log.Println("Connected to PostgreSQL successfully")

	return &DatabaseConfig{
		DB: db,
	}
}
