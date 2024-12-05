package models


type Recruiter struct {
	ID                 int       `db:"recruiter_id"`
	CompanyName        string    `db:"company_name"`
	CompanyDescription string    `db:"company_description"`
	CompanyWebsite     string    `db:"company_website"`
	VerifiedStatus     bool      `db:"verified_status"`
	UserID             int       `db:"recruiter_id"`
}
