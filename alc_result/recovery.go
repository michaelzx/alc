package alc_result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RecoveryMiddleware(unknownErrorHandler func(err interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch e := err.(type) {
				case *CommonErr:
					c.JSON(http.StatusOK, e)
					c.AbortWithStatus(http.StatusOK)
				case *UnauthorizedErr:
					c.JSON(http.StatusUnauthorized, e)
					c.AbortWithStatus(http.StatusUnauthorized)
				case *ForbiddenErr:
					c.JSON(http.StatusForbidden, e)
					c.AbortWithStatus(http.StatusForbidden)
				case *NotFoundErr:
					c.JSON(http.StatusNotFound, e)
					c.AbortWithStatus(http.StatusNotFound)
				default:
					if unknownErrorHandler != nil {
						unknownErrorHandler(err)
					}
					c.JSON(http.StatusInternalServerError, &CommonErr{http.StatusInternalServerError, "服务器繁忙", nil})
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()
		c.Next()
	}
}
