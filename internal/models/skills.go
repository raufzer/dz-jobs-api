package models

import (
	"github.com/google/uuid"
)

type CandidateSkills struct {
	ID    uuid.UUID `db:"candidate_id"`
	Skill string    `db:"skill"`
}
