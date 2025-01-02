package v1

import (
	"dz-jobs-api/internal/controllers"
	"dz-jobs-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SystemRoutes(rg *gin.RouterGroup, systemController *controllers.SystemController) {
	rg.GET("/", systemController.DefaultRoute)
	rg.GET("/version", systemController.GetAPIVersion)
	rg.GET("/health", systemController.GetHealth)
	rg.Use(middlewares.MetricsMiddleware())
	rg.GET("/metrics", systemController.GetMetrics)
}
