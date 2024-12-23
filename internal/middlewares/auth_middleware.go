package middlewares

import (
	"dz-jobs-api/config"
	"dz-jobs-api/pkg/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(config *config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {

		accessToken, err := c.Cookie("access_token")
		if err != nil {
			c.Error(utils.NewCustomError(http.StatusUnauthorized, "No access token found"))
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(accessToken, config.AccessTokenSecret, "access")
		if err != nil {
			c.Error(utils.NewCustomError(http.StatusUnauthorized, "Invalid or expired access token"))
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)
		c.Set("role", claims.Role)
		if claims.Role == "candidate" {
			c.Set("candidate_id", claims.UserID)
		} else if claims.Role == "recruiter" {
			c.Set("recruiter_id", claims.UserID)
		}
		c.Set("purpose", claims.Purpose)
		c.Next()
	}
}
