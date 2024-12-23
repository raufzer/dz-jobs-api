package request



type PostNewJobRequest struct {
	Title          string     `json:"title" binding:"required"`
	Description    string     `json:"description" binding:"required"`
	Location       string     `json:"location,omitempty"`
	SalaryRange    string     `json:"salary_range,omitempty"`
	RequiredSkills string     `json:"required_skills,omitempty"`
	Status         string     `json:"status" binding:"required"`
}

type EditJobRequest struct {
	Title          string     `json:"title,omitempty"`
	Description    string     `json:"description,omitempty"`
	Location       string     `json:"location,omitempty"`
	SalaryRange    string     `json:"salary_range,omitempty"`
	RequiredSkills string     `json:"required_skills,omitempty"`
	Status         string     `json:"status,omitempty"`
}

type JobFilters struct {
    Status          string   `json:"status"`
    Location        string   `json:"location"`
    SalaryRangeMin  float64  `json:"salary_range_min"`
    SalaryRangeMax  float64  `json:"salary_range_max"`
    Skills          []string `json:"skills"`
    Keyword         string   `json:"keyword"`
}