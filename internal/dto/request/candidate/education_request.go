package candidate

import "time"

type AddEducationRequest struct {
	Degree      string    `json:"degree" binding:"required"`
	Institution string    `json:"institution" binding:"required"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
}

type UpdateEducationRequest struct {
	Degree      string    `json:"degree"`
	Institution string    `json:"institution"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
}
