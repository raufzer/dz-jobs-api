package candidate

import (
	"time"


)

type AddExperienceRequest struct {
	JobTitle    string    `json:"job_title" binding:"required"`
	Company     string    `json:"company" binding:"required"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
}

type UpdateExperienceRequest struct {
	JobTitle    string    `json:"job_title"`
	Company     string    `json:"company"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
}
