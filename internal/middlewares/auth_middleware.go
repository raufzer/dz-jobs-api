package middlewares

import (
	"dz-jobs-api/config"
	"dz-jobs-api/pkg/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(config *config.AppConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		accessToken, err := ctx.Cookie("access_token")
		if err != nil {
			ctx.Error(utils.NewCustomError(http.StatusUnauthorized, "No access token found"))
			ctx.Abort()
			return
		}

		claims, err := utils.ValidateToken(accessToken, config.AccessTokenSecret, "access")
		if err != nil {
			ctx.Error(utils.NewCustomError(http.StatusUnauthorized, "Invalid or expired access token"))
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Set("role", claims.Role)
		ctx.Set("purpose", claims.Purpose)
		if claims.Role == "candidate" {
			ctx.Set("candidate_id", claims.UserID)
		} else if claims.Role == "recruiter" {
			ctx.Set("recruiter_id", claims.UserID)
		}
		ctx.Next()
	}
}
