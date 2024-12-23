package candidate

import (
	models "dz-jobs-api/internal/models/candidate"

	"github.com/google/uuid"
)

type ExperienceResponse struct {
	ExperienceID uuid.UUID `json:"experience_id"`
	CandidateID  uuid.UUID `json:"candidate_id"`
	JobTitle     string    `json:"job_title"`
	Company      string    `json:"company"`
	StartDate    string    `json:"start_date"`
	EndDate      string    `json:"end_date"`
	Description  string    `json:"description"`
}

func ToExperienceResponse(experience *models.CandidateExperience) ExperienceResponse {
	return ExperienceResponse{
		ExperienceID: experience.ExperienceID,
		CandidateID:  experience.CandidateID,
		JobTitle:     experience.JobTitle,
		Company:      experience.Company,
		StartDate:    experience.StartDate,
		EndDate:      experience.EndDate,
		Description:  experience.Description,
	}
}

type ExperiencesResponseData struct {
	Total       int                `json:"total"`
	Experiences []ExperienceResponse `json:"experiences"`
}

func ToExperiencesResponse(experiences []models.CandidateExperience) ExperiencesResponseData {
	var experienceResponses []ExperienceResponse
	for _, exp := range experiences {
		experienceResponses = append(experienceResponses, ToExperienceResponse(&exp))
	}
	return ExperiencesResponseData{
		Total:       len(experiences),
		Experiences: experienceResponses,
	}
}
