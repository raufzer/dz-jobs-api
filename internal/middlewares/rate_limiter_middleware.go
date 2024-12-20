package middlewares

import (
	"dz-jobs-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiter(limit int, burst int) gin.HandlerFunc {

	limiter := rate.NewLimiter(rate.Limit(limit), burst)

	return func(ctx *gin.Context) {

		if !limiter.Allow() {

			ctx.JSON(http.StatusTooManyRequests, response.Response{
				Code:    http.StatusTooManyRequests,
				Status:  "Too Many Requests",
				Message: "Rate limit exceeded. Please try again later.a",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
