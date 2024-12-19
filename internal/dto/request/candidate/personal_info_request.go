package candidate

import "github.com/google/uuid"

type UpdateCandidatePersonalInfoRequest struct {
	CandidateID uuid.UUID `json:"candidate_id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
}
