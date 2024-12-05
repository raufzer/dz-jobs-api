package request

type CreateRecruiterRequest struct {
	CompanyName        string `json:"companyName" binding:"required"`
	CompanyDescription string `json:"companyDescription" binding:"required"`
	CompanyWebsite     string `json:"companyWebsite" binding:"required,url"`
	VerifiedStatus     bool   `json:"verifiedStatus"`
}

type UpdateRecruiterRequest struct {
	CompanyName        string `json:"companyName,omitempty" validate:"omitempty,min=3,max=100"`
	CompanyDescription string `json:"companyDescription,omitempty"`
	CompanyWebsite     string `json:"companyWebsite,omitempty" validate:"omitempty,url"`
	VerifiedStatus     bool   `json:"verifiedStatus,omitempty"`
}
