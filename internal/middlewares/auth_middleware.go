package middlewares

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/helpers"
	"dz-jobs-api/pkg/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(config *config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {

		accessToken, err := c.Cookie("access_token")
		if err != nil {
			helpers.NewCustomError(http.StatusUnauthorized, "No access token found")
			c.Abort()
			return
		}

		userID, err := utils.ValidateToken(accessToken, config.AccessTokenSecret)
		if err != nil {
			helpers.NewCustomError(http.StatusUnauthorized, "Invalid or expired access token")
			c.Abort()
			return
		}

		c.Set("userID", userID)

		c.Next()
	}
}
