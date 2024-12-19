package candidate

import (


	"github.com/google/uuid"
)

type CandidateSkills struct {
	CandidateID uuid.UUID `db:"candidate_id"`
	Skill       string    `db:"skill"`
}