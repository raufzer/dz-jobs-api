package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCORS(domain string) gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{domain}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Authorization", "Content-Type")
	return cors.New(corsConfig)
}
