package utils

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetAuthCookie(ctx *gin.Context, tokenName string, token string, maxAge time.Duration, domain string, isProduction bool) {
	cookie := &http.Cookie{
		Name:     tokenName,
		Value:    token,
		Path:     "/",
		Domain:   domain,
		MaxAge:   int(maxAge.Seconds()),
		Secure:   isProduction,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}

	http.SetCookie(ctx.Writer, cookie)
}
