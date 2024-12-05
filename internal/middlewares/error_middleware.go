package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"dz-jobs-api/helpers"
	"dz-jobs-api/internal/dto/response"

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
		switch {
		case helpers.IsErrorType(e.Err, helpers.ErrEmailAlreadyExists):
			ctx.JSON(http.StatusConflict, response.Response{
				Code:    http.StatusConflict,
				Status:  "Conflict",
				Message: "Email already exists",
			})
			ctx.Abort()
			return

		case helpers.IsErrorType(e.Err, helpers.ErrInvalidUserData):
			ctx.JSON(http.StatusBadRequest, response.Response{
				Code:    http.StatusBadRequest,
				Status:  "Bad Request",
				Message: "Invalid user data",
				Data:    e.Error(),
			})
			ctx.Abort()
			return

		case helpers.IsErrorType(e.Err, helpers.ErrUserCreationFailed):
			log.Printf("User creation error: %v", e.Err)
			ctx.JSON(http.StatusInternalServerError, response.Response{
				Code:    http.StatusInternalServerError,
				Status:  "Internal Server Error",
				Message: "Failed to create user",
				Data:    e.Error(),
			})
			ctx.Abort()
			return

		case isValidationError(e.Err):
			handleValidationError(ctx, e.Err)
			ctx.Abort()
			return

		case helpers.IsErrorType(e.Err, helpers.ErrInvalidCredentials):
			ctx.JSON(http.StatusUnauthorized, response.Response{
				Code:    http.StatusUnauthorized,
				Status:  "Unauthorized",
				Message: "Invalid email or password",
			})
			ctx.Abort()
			return

		case helpers.IsErrorType(e.Err, helpers.ErrTokenGeneration):
			ctx.JSON(http.StatusInternalServerError, response.Response{
				Code:    http.StatusInternalServerError,
				Status:  "Internal Server Error",
				Message: "Failed to generate authentication token",
			})
			ctx.Abort()
			return

		case helpers.IsErrorType(e.Err, helpers.ErrInvalidUserData):
			ctx.JSON(http.StatusBadRequest, response.Response{
				Code:    http.StatusBadRequest,
				Status:  "Bad Request",
				Message: "Invalid login data",
				Data:    e.Error(),
			})
			ctx.Abort()
			return
		case helpers.IsErrorType(e.Err, helpers.ErrUserCreationFailed):
			ctx.JSON(http.StatusInternalServerError, response.Response{
				Code:    http.StatusInternalServerError,
				Status:  "Internal Server Error",
				Message: "Failed to create user",
			})
			ctx.Abort()
			return

		case helpers.IsErrorType(e.Err, helpers.ErrUserNotFound):
			ctx.JSON(http.StatusNotFound, response.Response{
				Code:    http.StatusNotFound,
				Status:  "Not Found",
				Message: "User not found",
			})
			ctx.Abort()
			return

		case helpers.IsErrorType(e.Err, helpers.ErrInvalidUserData):
			ctx.JSON(http.StatusBadRequest, response.Response{
				Code:    http.StatusBadRequest,
				Status:  "Bad Request",
				Message: "Invalid user data",
				Data:    e.Error(),
			})
			ctx.Abort()
			return

		default:
			log.Printf("Unhandled error: %v", e.Err)
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

func isValidationError(err error) bool {
	_, ok := err.(validator.ValidationErrors)
	return ok
}

func handleValidationError(ctx *gin.Context, err error) {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return
	}

	var errorDetails []string
	for _, e := range validationErrors {
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
