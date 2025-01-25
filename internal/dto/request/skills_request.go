package request

type AddSkillRequest struct {
	Skill string `json:"skill" binding:"required"`
}
