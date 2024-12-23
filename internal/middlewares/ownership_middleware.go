package middlewares

import (
	"dz-jobs-api/pkg/utils"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CandidateOwnershipMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		candidateID := c.Param("candidate_id")

		authCandidateID, exists := c.Get("candidate_id")
		if !exists || candidateID != authCandidateID {
			c.Error(utils.NewCustomError(http.StatusForbidden, "Forbidden: You do not own this resource"))
			c.Abort()
			return
		}

		c.Next()
	}
}

func RecruiterOwnershipMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		recruiterID := c.Param("recruiter_id")

		authRecruiterID, exists := c.Get("recruiter_id")
		if !exists || recruiterID != authRecruiterID {
			c.Error(utils.NewCustomError(http.StatusForbidden, "Forbidden: You do not own this resource"))
			c.Abort()
			return
		}

		c.Next()
	}
}