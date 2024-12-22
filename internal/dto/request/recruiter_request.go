package request

type CreateRecruiterRequest struct {
	CompanyName        string `form:"company_name" binding:"required"`
	CompanyDescription string `form:"company_description" binding:"required"`
	CompanyWebsite     string `form:"company_website" binding:"required,url"`
	CompanyLocation    string `form:"company_location" binding:"required"`
	CompanyContact     string `form:"company_contact" binding:"required"`
	SocialLinks        string `form:"social_links" binding:"required"`
	VerifiedStatus     bool   `form:"verified_status" binding:"required"`
}

type UpdateRecruiterRequest struct {
	CompanyName        string `form:"company_name,omitempty" validate:"omitempty,min=3,max=150"`
	CompanyDescription string `form:"company_description,omitempty" validate:"omitempty"`
	CompanyWebsite     string `form:"company_website,omitempty" validate:"omitempty,url"`
	CompanyLocation    string `form:"company_location,omitempty" validate:"omitempty"`
	CompanyContact     string `form:"company_contact,omitempty" validate:"omitempty"`
	SocialLinks        string `form:"social_links" binding:"required"`
	VerifiedStatus     bool   `form:"verified_status,omitempty" validate:"omitempty"`
}
