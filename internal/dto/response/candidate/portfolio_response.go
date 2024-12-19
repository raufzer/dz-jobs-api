package candidate

import (
	models "dz-jobs-api/internal/models/candidate"

	"github.com/google/uuid"
)

type PortfolioResponse struct {
	ProjectID   uuid.UUID `json:"project_id"`
	CandidateID uuid.UUID `json:"candidate_id"`
	ProjectName string    `json:"project_name"`
	ProjectLink string    `json:"project_link"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
}

func ToPortfolioResponse(portfolio *models.CandidatePortfolio) PortfolioResponse {
	return PortfolioResponse{
		ProjectID:   portfolio.ProjectID,
		CandidateID: portfolio.CandidateID,
		ProjectName: portfolio.ProjectName,
		ProjectLink: portfolio.ProjectLink,
		Category:    portfolio.Category,
		Description: portfolio.Description,
	}
}
