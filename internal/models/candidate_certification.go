package models

import (
	"time"

	"github.com/google/uuid"
)
type CandidateCertification struct {
	CertificationID   uuid.UUID `db:"certification_id"`
	CandidateID       uuid.UUID `db:"candidate_id"`
	CertificationName string    `db:"certification_name"`
	IssuedBy          string    `db:"issued_by"`
	IssueDate         time.Time `db:"issue_date"`
	ExpirationDate    time.Time `db:"expiration_date"`
}