package middlewares

import (
	"dz-jobs-api/pkg/utils"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CandidateOwnershipMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		candidateID := ctx.Param("candidate_id")

		authCandidateID, exists := ctx.Get("candidate_id")
		if !exists || candidateID != authCandidateID {
			ctx.Error(utils.NewCustomError(http.StatusForbidden, "Forbidden: You do not own this resource"))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func RecruiterOwnershipMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		recruiterID := ctx.Param("recruiter_id")

		authRecruiterID, exists := ctx.Get("recruiter_id")
		if !exists || recruiterID != authRecruiterID {
			ctx.Error(utils.NewCustomError(http.StatusForbidden, "Forbidden: You do not own this resource"))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}