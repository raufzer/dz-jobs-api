package helpers

import (
	"github.com/gin-gonic/gin"
	"time"
)

func SetAuthCookie(ctx *gin.Context, token string, maxAge time.Duration, domain string) {
	ctx.SetCookie(
		"token",
		token,
		int(maxAge.Seconds()),
		"/",
		domain,
		false,
		true,
	)
}
