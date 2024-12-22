package models

import (
	"github.com/google/uuid"
)

type Recruiter struct {
	RecruiterID        uuid.UUID `db:"recruiter_id"`
	CompanyName        string    `db:"company_name"`
	CompanyLogo        string    `db:"company_logo"`
	CompanyDescription string    `db:"company_description"`
	CompanyWebsite     string    `db:"company_website"`
	CompanyLocation    string    `db:"company_location"`
	CompanyContact     string    `db:"company_contact"`
	SocialLinks        string    `db:"social_links"`
	VerifiedStatus     bool      `db:"verified_status"`
}
