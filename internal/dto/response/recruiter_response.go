package response

import (
	"dz-jobs-api/internal/models"
)

type RecruiterResponse struct {
	ID                 int    `json:"id"`
	CompanyName        string `json:"companyName"`
	CompanyDescription string `json:"companyDescription"`
	CompanyWebsite     string `json:"companyWebsite"`
	VerifiedStatus     bool   `json:"verifiedStatus"`
}

func ToRecruiterResponse(recruiter *models.Recruiter) RecruiterResponse {
	return RecruiterResponse{
		ID:                 recruiter.ID,
		CompanyName:        recruiter.CompanyName,
		CompanyDescription: recruiter.CompanyDescription,
		CompanyWebsite:     recruiter.CompanyWebsite,
		VerifiedStatus:     recruiter.VerifiedStatus,
	}
}
