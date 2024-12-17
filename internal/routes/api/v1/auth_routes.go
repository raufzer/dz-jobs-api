package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup, authController *controllers.AuthController) {
	authRoute := rg.Group("/auth")
	authRoute.POST("/register", authController.Register)
	authRoute.POST("/login", authController.Login)
	authRoute.POST("/logout", authController.Logout)
	authRoute.POST("/refresh-token", authController.RefreshToken)
	authRoute.POST("/send-reset-otp", authController.SendResetOTP)
	authRoute.POST("/verify-otp", authController.VerifyOTP)
	authRoute.POST("/reset-password", authController.ResetPassword)
	authRoute.GET("/google/connect", authController.GoogleConnect)
	authRoute.GET("/google/callback", authController.GoogleCallbackConnect)

}
