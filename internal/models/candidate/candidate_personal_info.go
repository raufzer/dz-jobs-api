package candidate

import (

	"github.com/google/uuid"
)

type CandidatePersonalInfo struct {
	CandidateID uuid.UUID `db:"candidate_id"`
	Name        string    `db:"name"`
	Email       string    `db:"email"`
	Phone       string    `db:"phone"`
	Address     string    `db:"address"`
}