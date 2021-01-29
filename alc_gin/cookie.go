package alc_gin

import (
	"github.com/gin-gonic/gin"
)

func CookieRemove(c *gin.Context, cookieName string) {
	c.SetCookie(cookieName, "", 0, "", "", false, true)
}

func CookieAdd(c *gin.Context, cookieName string, cookieValue string, maxAge int) {
	c.SetCookie(cookieName, cookieValue, maxAge, "", "", false, true)
}
