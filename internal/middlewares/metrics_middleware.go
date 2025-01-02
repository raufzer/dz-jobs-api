package middlewares

import (
	"dz-jobs-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

// MetricsMiddleware tracks requests and errors
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Increment request count
		utils.IncrementRequest()

		c.Next()

		// Increment error count if status is >= 400
		if c.Writer.Status() >= 400 {
			utils.IncrementError()
		}
	}
}
