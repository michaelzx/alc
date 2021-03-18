package alc_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelzx/alc/alc_errs"
	"github.com/michaelzx/alc/alc_logger"
	"go.uber.org/zap"
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
				alc_logger.Error("get error from recover", zap.Any("unknown error", err))
				c.AbortWithStatus(unknown.Status)
				return
			}
		}()
		c.Next()
	}
}
