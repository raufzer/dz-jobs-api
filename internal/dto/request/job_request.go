package request

type PostNewJobRequest struct {
	Title          string `json:"title" validate:"required"`
	Description    string `json:"description" validate:"required"`
	Location       string `json:"location,omitempty"`
	SalaryRange    string `json:"salary_range,omitempty" validate:"omitempty,jobSalaryRange"`
	RequiredSkills string `json:"required_skills,omitempty"`
	Status         string `json:"status" validate:"required,oneof=open closed"`
	JobType        string `json:"job_type" validate:"required,oneof=full-time part-time freelance remote"`
}

type EditJobRequest struct {
	Title          string `json:"title,omitempty"`
	Description    string `json:"description,omitempty"`
	Location       string `json:"location,omitempty"`
	SalaryRange    string `json:"salary_range,omitempty"`
	RequiredSkills string `json:"required_skills,omitempty"`
	Status         string `json:"status,omitempty" binding:"omitempty,oneof=open closed"`
	JobType        string `json:"job_type,omitempty" binding:"omitempty,oneof=full-time part-time freelance remote"`
}

type JobFilters struct {
	Status         string   `form:"status"`
	Location       string   `form:"location"`
	SalaryRangeMin float64  `form:"min_salary"`
	SalaryRangeMax float64  `form:"max_salary"`
	RequiredSkills []string `form:"required_skills"`
	Keyword        string   `form:"keyword"`
	JobType        string   `form:"job_type"`
}
