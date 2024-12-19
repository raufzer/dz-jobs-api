package candidate

import "github.com/google/uuid"

type AddSkillRequest struct {
	CandidateID uuid.UUID `json:"candidate_id" binding:"required"`
	Skill       string    `json:"skill" binding:"required"`
}

type RemoveSkillRequest struct {
	CandidateID uuid.UUID `json:"candidate_id" binding:"required"`
	Skill       string    `json:"skill" binding:"required"`
}
