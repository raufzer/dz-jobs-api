package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RecruiterJobRoutes(rg *gin.RouterGroup, jobController *controllers.JobController) {
	jobs := rg.Group("/jobs")
	jobs.POST("/", jobController.PostNewJob)
	jobs.GET("/:job_id", jobController.GetJobDetails)
	jobs.GET("/", jobController.GetJobListingsByStatus)
	jobs.PUT("/:job_id", jobController.EditJob)
	jobs.PUT("/:job_id/deactivate", jobController.DeactivateJob)
	jobs.PUT("/:job_id/repost", jobController.RepostJob)
	jobs.DELETE("/:job_id", jobController.DeleteJob)
}

func JobRoutes(rg *gin.RouterGroup, jobController *controllers.JobController) {
	jobs := rg.Group("/jobs")
	jobs.GET("/", jobController.GetAllJobs)
	jobs.GET("/search", jobController.SearchJobs)
	jobs.GET("/:job_id", jobController.GetJobDetailsPublic)
}
