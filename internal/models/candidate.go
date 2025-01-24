package models

import (
	"github.com/google/uuid"
)

type Candidate struct {
	ID             uuid.UUID `db:"candidate_id"`
	Resume         string    `db:"resume"`
	ProfilePicture string    `db:"profile_picture"`
}
