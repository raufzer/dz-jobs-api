package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`             // Primary key
	Name      string    `json:"name"`           // User's name
	Email     string    `json:"email"`          // User's email
	Password  string    `json:"-"`              // User's password (omit from JSON)
	Role      string    `json:"role,omitempty"` // Optional role
	CreatedAt time.Time `json:"created_at"`     // Record creation timestamp
	UpdatedAt time.Time `json:"updated_at"`     // Record update timestamp
}
