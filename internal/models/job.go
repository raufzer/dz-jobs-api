package models

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	JobID          int64     `db:"job_id"`
	Title          string    `db:"title"`
	Description    string    `db:"description"`
	Location       string    `db:"location,omitempty"`
	SalaryRange    string    `db:"salary_range,omitempty"`
	RequiredSkills string    `db:"required_skills,omitempty"`
	RecruiterID    uuid.UUID `db:"recruiter_id"`
	CreatedAt      time.Time `db:"created_at" default:"CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `db:"updated_at" default:"CURRENT_TIMESTAMP"`
    Status         string    `db:"status"`
    JobType        string    `db:"job_type"`
}
