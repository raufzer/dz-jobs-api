package candidate

import (
	"time"

	"github.com/google/uuid"
)

type AddExperienceRequest struct {
	CandidateID uuid.UUID `json:"candidate_id" binding:"required"`
	JobTitle    string    `json:"job_title" binding:"required"`
	Company     string    `json:"company" binding:"required"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
}

type UpdateExperienceRequest struct {
	CandidateID uuid.UUID `json:"candidate_id" binding:"required"`
	JobTitle    string    `json:"job_title"`
	Company     string    `json:"company"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
}
