package response

import (
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type SkillResponse struct {
	ID    uuid.UUID `json:"candidate_id"`
	Skill string    `json:"skill"`
}

func ToSkillResponse(skill *models.CandidateSkills) SkillResponse {
	return SkillResponse{
		ID:    skill.ID,
		Skill: skill.Skill,
	}
}

type SkillsResponseData struct {
	Total  int             `json:"total"`
	Skills []SkillResponse `json:"skills"`
}

func ToSkillsResponse(skills []models.CandidateSkills) SkillsResponseData {
	var skillResponses []SkillResponse
	for _, skill := range skills {
		skillResponses = append(skillResponses, ToSkillResponse(&skill))
	}
	return SkillsResponseData{
		Total:  len(skills),
		Skills: skillResponses,
	}
}
