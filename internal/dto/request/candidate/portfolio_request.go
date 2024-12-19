package candidate

import "github.com/google/uuid"

type AddPortfolioRequest struct {
	CandidateID uuid.UUID `json:"candidate_id" binding:"required"`
	ProjectName string    `json:"project_name" binding:"required"`
	ProjectLink string    `json:"project_link" binding:"required,url"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
}

type UpdatePortfolioRequest struct {
	CandidateID uuid.UUID `json:"candidate_id" binding:"required"`
	ProjectName string    `json:"project_name"`
	ProjectLink string    `json:"project_link"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
}
