package controllers

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/dto/response"
	"dz-jobs-api/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SystemController struct {
config *config.AppConfig
}

func NewSystemController(config *config.AppConfig) *SystemController {
	return &SystemController{
		config: config,
	}
}

// GetVersion provides API version and metadata
// @Summary Get API version and metadata
// @Description Returns metadata about the API, including version, environment, build details, and health status.
// @Tags Version
// @Produce json
// @Success 200 {object} response.VersionResponse
// @Router /version [get]
func (c *SystemController) GetAPIVersion(ctx *gin.Context) {
	loc, _ := time.LoadLocation("UTC")
	ctx.JSON(http.StatusOK, response.VersionResponse{
		APIVersion:       1,
		BuildVersion:     c.config.BuildVersion,
		CommitHash:       c.config.CommitHash,
		ReleaseDate:      time.Now().In(loc).Format("2006-01-02"),
		Environment:      c.config.Environment,
		Status:           "healthy",
		DocumentationURL: c.config.DocumentationURL,
		LastMigration:    c.config.LastMigration,
	})
}

// GetHealth provides API health status
// @Summary Get API health status
// @Description Returns the health status of the API.
// @Tags Health
// @Produce json
// @Success 200 {object} response.HealthResponse "API is healthy"
// @Router /health [get]
func (c *SystemController) GetHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.HealthResponse{
		Status: "healthy"})
}

// GetMetrics provides API metrics
// @Summary Get API metrics including uptime, request count, and error rate
// @Description Returns the metrics for the API including uptime, request count, and error rate.
// @Tags Metrics
// @Produce json
// @Success 200 {object} response.MetricsResponse "API metrics data"
// @Router /metrics [get]
func (c *SystemController)  GetMetrics(ctx *gin.Context) {
	uptime, requestCount, errorRate := utils.GetMetrics()
	ctx.JSON(http.StatusOK, response.MetricsResponse{
		Uptime:       uptime,
		RequestCount: requestCount,
		ErrorRate:    errorRate,
	})
}
