package request

type AddProjectRequest struct {
	ProjectName string `json:"project_name" binding:"required"`
	ProjectLink string `json:"project_link" binding:"required,url"`
	Category    string `json:"category"`
	Description string `json:"description"`
}
