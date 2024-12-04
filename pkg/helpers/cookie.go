package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetAuthCookie(ctx *gin.Context, token string, maxAge time.Duration, domain string) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Domain:   domain,
		MaxAge:   int(maxAge.Seconds()),
		Secure:   true, // Only sent over HTTPS
		HttpOnly: true, // Cannot be accessed by JavaScript
		SameSite: http.SameSiteNoneMode,
	}

	// Set the cookie on the response
	http.SetCookie(ctx.Writer, cookie)
}
