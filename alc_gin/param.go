package alc_gin

import (
	"github.com/gin-gonic/gin"
)

func ParamToBool(c *gin.Context, paramName string) *bool {
	if c.Param(paramName) == "true" {
		boolTrue := true
		return &boolTrue
	}
	if c.Param(paramName) == "false" {
		boolTrue := false
		return &boolTrue
	}
	return nil
}

// TODO 待定
// func ParamInt64(c *gin.Context, paramName string) *int64 {
// 	str := c.Param(paramName)
// 	if str == "" {
// 		return nil
// 	}
// 	i64, err := strconv.ParseInt(str, 10, 64)
// 	if err != nil {
// 		return nil
// 	}
// 	return &i64
// }
// func ParamInt64Default(c *gin.Context, paramName string, defaultValue int64) int64 {
// 	str := c.Param(paramName)
// 	if str == "" {
// 		return defaultValue
// 	}
// 	i64, err := strconv.ParseInt(str, 10, 64)
// 	if err != nil {
// 		return defaultValue
// 	}
// 	return i64
// }
