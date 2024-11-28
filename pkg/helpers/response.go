package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"dz-jobs-api/data/response"
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

