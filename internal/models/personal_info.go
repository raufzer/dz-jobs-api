package models

import (
	"github.com/google/uuid"
)

type CandidatePersonalInfo struct {
	ID          uuid.UUID `db:"candidate_id"`
	Name        string    `db:"name"`
	Email       string    `db:"email"`
	Phone       string    `db:"phone"`
	Address     string    `db:"address"`
	DateOfBirth string    `db:"date_of_birth"`
	Gender      string    `db:"gender"`
	Bio         string    `db:"bio"`
}
