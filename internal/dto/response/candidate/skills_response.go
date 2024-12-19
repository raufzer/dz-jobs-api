package candidate

import (
	models "dz-jobs-api/internal/models/candidate"

	"github.com/google/uuid"
)

type SkillResponse struct {
	CandidateID uuid.UUID `json:"candidate_id"`
	Skill       string    `json:"skill"`
}


func ToSkillsResponse(skill *models.CandidateSkills) SkillResponse {
    return SkillResponse{
        CandidateID: skill.CandidateID,
        Skill:       skill.Skill,
    }
}
