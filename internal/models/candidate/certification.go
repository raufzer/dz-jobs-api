package candidate

import (
	"github.com/google/uuid"
)

type CandidateCertification struct {
	CertificationID   uuid.UUID `db:"certification_id"`
	CandidateID       uuid.UUID `db:"candidate_id"`
	CertificationName string    `db:"certification_name"`
	IssuedBy          string    `db:"issued_by"`
	IssueDate         string    `db:"issue_date"`
	ExpirationDate    string    `db:"expiration_date"`
}
