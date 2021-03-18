package alc_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelzx/alc/alc_lang"
)

func GetLangTag(gc *gin.Context) alc_lang.Tag {
	tag, exist := gc.Get(alc_lang.GinContextKey)
	if !exist {
		return alc_lang.None
	}
	v, ok := tag.(alc_lang.Tag)
	if !ok {
		return alc_lang.None
	}
	return v
}
