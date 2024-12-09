package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"dz-jobs-api/internal/dto/response"
	"dz-jobs-api/internal/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic recovered: %v", r)
				ctx.JSON(http.StatusInternalServerError, response.Response{
					Code:    http.StatusInternalServerError,
					Status:  "Internal Server Error",
					Message: "An unexpected error occurred",
				})
				ctx.Abort()
			}
		}()

		ctx.Next()

		if len(ctx.Errors) > 0 {
			handleErrors(ctx)
		}
	}
}

func handleErrors(ctx *gin.Context) {
	for _, e := range ctx.Errors {
		switch err := e.Err.(type) {
		case *helpers.CustomError:

			ctx.JSON(err.StatusCode, response.Response{
				Code:    err.StatusCode,
				Status:  http.StatusText(err.StatusCode),
				Message: err.Message,
			})
			ctx.Abort()
			return

		case validator.ValidationErrors:

			handleValidationError(ctx, err)
			ctx.Abort()
			return

		default:

			log.Printf("Unhandled error: %v", err)
			ctx.JSON(http.StatusInternalServerError, response.Response{
				Code:    http.StatusInternalServerError,
				Status:  "Internal Server Error",
				Message: "An unexpected error occurred",
			})
			ctx.Abort()
			return
		}
	}
}

func handleValidationError(ctx *gin.Context, err validator.ValidationErrors) {
	var errorDetails []string
	for _, e := range err {
		errorDetails = append(errorDetails, fmt.Sprintf(
			"Field: %s, Error: %s, Value: %v",
			e.Field(),
			e.Tag(),
			e.Value(),
		))
	}

	ctx.JSON(http.StatusBadRequest, response.Response{
		Code:    http.StatusBadRequest,
		Status:  "Validation Error",
		Message: "Invalid input data",
		Data:    strings.Join(errorDetails, "; "),
	})
}
