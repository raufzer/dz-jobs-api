package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func JobRoutes(rg *gin.RouterGroup, jobController *controllers.JobController) {
	jobs := rg.Group("/jobs")
	{
		jobs.POST("/:recruiter_id", jobController.PostNewJob)
		jobs.GET("/:job_id", jobController.GetJobDetails)
		jobs.GET("/", jobController.GetJobListingsByStatus)
		jobs.PUT("/:job_id", jobController.EditJob)
		jobs.PUT("/:job_id/desactivate", jobController.DesactivateJob)
		jobs.PUT("/:job_id/repost", jobController.RepostJob)
		jobs.DELETE("/:job_id", jobController.DeleteJob)
	}
}
