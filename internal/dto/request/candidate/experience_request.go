package candidate



type AddExperienceRequest struct {
	JobTitle    string `json:"job_title" binding:"required"`
	Company     string `json:"company" binding:"required"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
}

type UpdateExperienceRequest struct {
	JobTitle    string `json:"job_title"`
	Company     string `json:"company"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
}
