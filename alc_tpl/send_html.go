package alc_tpl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendHtml4Gin(ctx *gin.Context, htmlStr string) {
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlStr))
}
