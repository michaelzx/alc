package alc_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelzx/alc/alc_errs"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if e, ok := err.(*alc_errs.BadRequest); ok {
					c.JSON(e.Status, e.BizErr)
					c.AbortWithStatus(e.Status)
					return
				}
				if e, ok := err.(*alc_errs.Unauthorized); ok {
					c.JSON(e.Status, e.BizErr)
					c.AbortWithStatus(e.Status)
					return
				}
				if e, ok := err.(*alc_errs.Forbidden); ok {
					c.JSON(e.Status, e.BizErr)
					c.AbortWithStatus(e.Status)
					return
				}
				if e, ok := err.(*alc_errs.NotFound); ok {
					c.JSON(e.Status, e.BizErr)
					c.AbortWithStatus(e.Status)
					return
				}
				unknown := alc_errs.NewUnknown("服务器繁忙")
				c.JSON(unknown.Status, unknown.BizErr)
				c.AbortWithStatus(unknown.Status)
				return
			}
		}()
		c.Next()
	}
}
