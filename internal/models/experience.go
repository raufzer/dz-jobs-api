package models

import (
	"github.com/google/uuid"
)

type CandidateExperience struct {
	ExperienceID uuid.UUID `db:"experience_id"`
	CandidateID  uuid.UUID `db:"candidate_id"`
	JobTitle     string    `db:"job_title"`
	Company      string    `db:"company"`
	StartDate    string    `db:"start_date"`
	EndDate      string    `db:"end_date"`
	Description  string    `db:"description"`
}
