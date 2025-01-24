package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RecruiterJobRoutes(rg *gin.RouterGroup, jobController *controllers.JobController) {
	jobs := rg.Group("/jobs")
	jobs.POST("/", jobController.PostNewJob)
	jobs.GET("/:jobId", jobController.GetJobDetails)
	jobs.GET("/", jobController.GetJobListingsByStatus)
	jobs.PUT("/:jobId", jobController.EditJob)
	jobs.PUT("/:jobId/deactivate", jobController.DeactivateJob)
	jobs.PUT("/:jobId/repost", jobController.RepostJob)
	jobs.DELETE("/:jobId", jobController.DeleteJob)
}

func JobRoutes(rg *gin.RouterGroup, jobController *controllers.JobController) {
	jobs := rg.Group("/jobs")
	jobs.GET("/", jobController.GetAllJobs)
	jobs.GET("/search", jobController.SearchJobs)
	jobs.GET("/:jobId", jobController.GetJobDetailsPublic)
}
