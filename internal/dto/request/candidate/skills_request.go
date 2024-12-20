package candidate



type AddSkillRequest struct {
	Skill       string    `json:"skill" binding:"required"`
}

type RemoveSkillRequest struct {
	Skill       string    `json:"skill" binding:"required"`
}
