package response

import (
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type EducationResponse struct {
	EducationID uuid.UUID `json:"education_id"`
	CandidateID uuid.UUID `json:"candidate_id"`
	Degree      string    `json:"degree"`
	Institution string    `json:"institution"`
	StartDate   string    `json:"start_date"`
	EndDate     string    `json:"end_date"`
	Description string    `json:"description"`
}

func ToEducationResponse(education *models.CandidateEducation) EducationResponse {
	return EducationResponse{
		EducationID: education.EducationID,
		CandidateID: education.CandidateID,
		Degree:      education.Degree,
		Institution: education.Institution,
		StartDate:   education.StartDate,
		EndDate:     education.EndDate,
		Description: education.Description,
	}
}

type EducationsResponseData struct {
	Total     int                `json:"total"`
	Educations []EducationResponse `json:"educations"`
}

func ToEducationsResponse(educations []models.CandidateEducation) EducationsResponseData {
	var educationResponses []EducationResponse
	for _, edu := range educations {
		educationResponses = append(educationResponses, ToEducationResponse(&edu))
	}
	return EducationsResponseData{
		Total:      len(educations),
		Educations: educationResponses,
	}
}
