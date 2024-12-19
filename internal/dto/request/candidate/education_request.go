package candidate

import (
	"time"

	"github.com/google/uuid"
)

type AddEducationRequest struct {
	CandidateID uuid.UUID `json:"candidate_id" binding:"required"`
	Degree      string    `json:"degree" binding:"required"`
	Institution string    `json:"institution" binding:"required"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
}

type UpdateEducationRequest struct {
	CandidateID uuid.UUID `json:"candidate_id"`
	Degree      string    `json:"degree"`
	Institution string    `json:"institution"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
}
