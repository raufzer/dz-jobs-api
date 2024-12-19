package candidate

import (
	"time"

	"github.com/google/uuid"
)

type CandidateExperience struct {
	ExperienceID uuid.UUID `db:"experience_id"`
	CandidateID  uuid.UUID `db:"candidate_id"`
	JobTitle     string    `db:"job_title"`
	Company      string    `db:"company"`
	StartDate    time.Time `db:"start_date"`
	EndDate      time.Time `db:"end_date"`
	Description  string    `db:"description"`
}
