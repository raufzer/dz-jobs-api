package candidate

import (
	"time"
)

type AddCertificationRequest struct {
	CertificationName string    `json:"certification_name" binding:"required"`
	IssuedBy          string    `json:"issued_by" binding:"required"`
	IssueDate         time.Time `json:"issue_date" binding:"required"`
	ExpirationDate    time.Time `json:"expiration_date"`
}

type UpdateCertificationRequest struct {
	CertificationName string    `json:"certification_name"`
	IssuedBy          string    `json:"issued_by"`
	IssueDate         time.Time `json:"issue_date"`
	ExpirationDate    time.Time `json:"expiration_date"`
}
