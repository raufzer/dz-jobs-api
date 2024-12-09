package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetAuthCookie(ctx *gin.Context, token string, maxAge time.Duration, domain string, isProduction bool) {
	cookie := &http.Cookie{
		Name:     "token",
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
