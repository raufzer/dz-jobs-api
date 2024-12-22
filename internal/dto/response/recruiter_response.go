package response

import (
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type RecruiterResponse struct {
	RecruiterID        uuid.UUID `json:"recruiter_id"`
	CompanyName        string    `json:"company_name"`
	CompanyLogo        string    `json:"company_logo"`
	CompanyDescription string    `json:"company_description"`
	CompanyWebsite     string    `json:"company_website"`
	CompanyLocation    string    `json:"company_location"`
	CompanyContact     string    `json:"company_contact"`
	SocialLinks        string    `json:"social_links"`
	VerifiedStatus     bool      `json:"verified_status"`
}

func ToRecruiterResponse(recruiter *models.Recruiter) RecruiterResponse {
	return RecruiterResponse{
		RecruiterID:        recruiter.RecruiterID,
		CompanyName:        recruiter.CompanyName,
		CompanyLogo:        recruiter.CompanyLogo,
		CompanyDescription: recruiter.CompanyDescription,
		CompanyWebsite:     recruiter.CompanyWebsite,
		CompanyLocation:    recruiter.CompanyLocation,
		CompanyContact:     recruiter.CompanyContact,
		SocialLinks:        recruiter.SocialLinks,
		VerifiedStatus:     recruiter.VerifiedStatus,
	}
}
