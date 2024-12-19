package middlewares

import (
	"dz-jobs-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// Create a custom middleware function for rate limiting
func RateLimiter(limit int, burst int) gin.HandlerFunc {
	// Create a rate limiter with specified limit and burst
	limiter := rate.NewLimiter(rate.Limit(limit), burst)

	return func(ctx *gin.Context) {
		// Check if the request can proceed
		if !limiter.Allow() {
			// If rate limit exceeded, return too many requests error
			ctx.JSON(http.StatusTooManyRequests, response.Response{
				Code:    http.StatusTooManyRequests,
				Status:  "Too Many Requests",
				Message: "Rate limit exceeded. Please try again later.a",
			})
			ctx.Abort()
			return
		}

		// If the request is allowed, continue to the next middleware/handler
		ctx.Next()
	}
}
