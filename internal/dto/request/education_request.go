package request



type AddEducationRequest struct {
	Degree      string `json:"degree" binding:"required"`
	Institution string `json:"institution" binding:"required"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
}

type UpdateEducationRequest struct {
	Degree      string `json:"degree"`
	Institution string `json:"institution"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
}
