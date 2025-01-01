package request



type AddExperienceRequest struct {
	JobTitle    string `json:"job_title" binding:"required"`
	Company     string `json:"company" binding:"required"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
}

