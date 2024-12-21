package candidate

type AddCertificationRequest struct {
	CertificationName string `json:"certification_name" binding:"required"`
	IssuedBy          string `json:"issued_by" binding:"required"`
	IssueDate         string `json:"issue_date" binding:"required"`
	ExpirationDate    string `json:"expiration_date"`
}

type UpdateCertificationRequest struct {
	CertificationName string `json:"certification_name"`
	IssuedBy          string `json:"issued_by"`
	IssueDate         string `json:"issue_date"`
	ExpirationDate    string `json:"expiration_date"`
}
