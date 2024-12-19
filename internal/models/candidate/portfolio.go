package candidate

import (

	"github.com/google/uuid"
)
type CandidatePortfolio struct {
	ProjectID   uuid.UUID `db:"project_id"`
	CandidateID uuid.UUID `db:"candidate_id"`
	ProjectName string    `db:"project_name"`
	ProjectLink string    `db:"project_link"`
	Category    string    `db:"category"`
	Description string    `db:"description"`
}