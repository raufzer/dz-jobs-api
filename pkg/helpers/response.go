package helpers

import (
	"net/http"

	"dz-jobs-api/data/response"
	"dz-jobs-api/internal/models"

	"github.com/gin-gonic/gin"
)

func RespondWithError(ctx *gin.Context, message string, statusCode int) {
	webResponse := response.Response{
		Code:    statusCode,
		Status:  http.StatusText(statusCode),
		Message: message,
	}
	ctx.JSON(statusCode, webResponse)
}

func RespondWithSuccess(ctx *gin.Context, message string, data interface{}) {
	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: message,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func ToUserResponse(user *models.User) response.UserResponse {
	return response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
