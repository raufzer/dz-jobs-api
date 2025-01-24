package models

import (
	"time"

	"github.com/google/uuid"
)

type Bookmark struct {
	ID          int64     `db:"bookmark_id"`
	CandidateID uuid.UUID `db:"candidate_id"`
	JobID       int64     `db:"job_id"`
	CreatedAt   time.Time `db:"created_at"`
}
