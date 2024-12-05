package models

import (
	"time"
)

type User struct {
	ID        int       `db:"userid"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
