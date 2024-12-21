package candidate

import (
	"github.com/google/uuid"
)

type CandidateEducation struct {
	EducationID uuid.UUID `db:"education_id"`
	CandidateID uuid.UUID `db:"candidate_id"`
	Degree      string    `db:"degree"`
	Institution string    `db:"institution"`
	StartDate   string    `db:"start_date"`
	EndDate     string    `db:"end_date"`
	Description string    `db:"description"`
}
