package alc_result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReplyOK(gc *gin.Context) {
	gc.JSON(http.StatusOK, &Result{
		Code: 0,
		Msg:  "success",
		Data: nil,
	})
}
func ReplyOkWithData(gc *gin.Context, data interface{}) {
	gc.JSON(http.StatusOK, &Result{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}
func ReplyFail(gc *gin.Context, code int, msg string) {
	gc.JSON(http.StatusOK, &Result{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
func ReplyFailWithData(gc *gin.Context, code int, msg string, data interface{}) {
	gc.JSON(http.StatusOK, &Result{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
