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
