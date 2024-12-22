package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCORS(frontDomain string, domain string) gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{frontDomain, domain}
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Authorization", "Content-Type")
	return cors.New(corsConfig)
}
