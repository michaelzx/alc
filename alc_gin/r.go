package alc_gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func RJson(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, v)
}
func RXmlString(c *gin.Context, xml string) {
	c.Data(http.StatusOK, "application/xml; charset=utf-8", []byte(xml))
}
func RString(c *gin.Context, str string) {
	c.String(http.StatusOK, str)
}
func RHtml(gc *gin.Context, htmlStr string) {
	gc.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlStr))
}

type Map map[string]interface{}

func RJsonSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Result{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}
func RJsonError(c *gin.Context, errCode int, errMsg string) {
	c.JSON(http.StatusOK, &Result{
		Code: errCode,
		Msg:  errMsg,
		Data: nil,
	})
}
func RJsonResult(c *gin.Context, errCode int, errMsg string, data interface{}) {
	c.JSON(http.StatusOK, &Result{
		Code: errCode,
		Msg:  errMsg,
		Data: data,
	})
}
