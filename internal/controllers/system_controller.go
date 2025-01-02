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

// DefaultRoute godoc
// @Summary Get the default route with API info
// @Description Returns a welcome message and useful API links, including version, health check, documentation, and metrics
// @Tags System - Default
// @Produce json
// @Success 200 {object} response.DefaultResponse
// @Router / [get]
func (c *SystemController) DefaultRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.DefaultResponse{
		Message:       "Welcome to the DZ Jobs API",
		Version:       c.config.VersionURL,
		Health:        c.config.HealthURL,
		Documentation: c.config.DocumentationURL,
		Metrics:       c.config.MetricsURL,
	})
}

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

func (c *SystemController) GetHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.HealthResponse{
		Status: "healthy"})
}

func (c *SystemController) GetMetrics(ctx *gin.Context) {
	uptime, requestCount, errorRate := utils.GetMetrics()
	ctx.JSON(http.StatusOK, response.MetricsResponse{
		Uptime:       uptime,
		RequestCount: requestCount,
		ErrorRate:    errorRate,
	})
}
