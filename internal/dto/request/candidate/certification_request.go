package candidate

import (
	"time"

	"github.com/google/uuid"
)

type AddCertificationRequest struct {
	CandidateID       uuid.UUID `json:"candidate_id" binding:"required"` 
	CertificationName string    `json:"certification_name" binding:"required"`
	IssuedBy          string    `json:"issued_by" binding:"required"`
	IssueDate         time.Time `json:"issue_date" binding:"required"`
	ExpirationDate    time.Time `json:"expiration_date"`
}

type UpdateCertificationRequest struct {
	CertificationID   uuid.UUID    `json:"certification_id" binding:"required"`
	CertificationName string    `json:"certification_name"`
	IssuedBy          string    `json:"issued_by"`
	IssueDate         time.Time `json:"issue_date"`
	ExpirationDate    time.Time `json:"expiration_date"`
}
