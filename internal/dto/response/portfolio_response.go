package response

import (
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type PortfolioResponse struct {
	ID          uuid.UUID `json:"project_id"`
	CandidateID uuid.UUID `json:"candidate_id"`
	ProjectName string    `json:"project_name"`
	ProjectLink string    `json:"project_link"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
}

func ToPortfolioResponse(portfolio *models.CandidatePortfolio) PortfolioResponse {
	return PortfolioResponse{
		ID:          portfolio.ID,
		CandidateID: portfolio.CandidateID,
		ProjectName: portfolio.ProjectName,
		ProjectLink: portfolio.ProjectLink,
		Category:    portfolio.Category,
		Description: portfolio.Description,
	}
}

type PortfoliosResponseData struct {
	Total    int                 `json:"total"`
	Projects []PortfolioResponse `json:"projects"`
}

func ToPortfoliosResponse(projects []models.CandidatePortfolio) PortfoliosResponseData {
	var portfolioResponses []PortfolioResponse
	for _, project := range projects {
		portfolioResponses = append(portfolioResponses, ToPortfolioResponse(&project))
	}
	return PortfoliosResponseData{
		Total:    len(projects),
		Projects: portfolioResponses,
	}
}
