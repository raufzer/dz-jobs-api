package middlewares

import (
	"dz-jobs-api/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		userRole, exists := c.Get("role")
		if !exists {
			c.Error(utils.NewCustomError(http.StatusUnauthorized, "Unauthorized: No role found"))
			c.Abort()
			return
		}

		roleStr, ok := userRole.(string)
		if !ok {
			c.Error(utils.NewCustomError(http.StatusInternalServerError, "Invalid role format"))
			c.Abort()
			return
		}

		for _, role := range allowedRoles {
			if strings.EqualFold(role, roleStr) {

				c.Next()
				return
			}
		}

		c.Error(utils.NewCustomError(http.StatusForbidden, "Forbidden: Access denied"))
		c.Abort()
	}
}

