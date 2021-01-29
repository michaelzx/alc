package alc_gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
