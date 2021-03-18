package alc_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/michaelzx/alc/alc_errs"
	"github.com/michaelzx/alc/alc_lang"
	"github.com/michaelzx/alc/alc_validator"
	"reflect"
	"strings"
)

func CheckDTO(context *gin.Context, reqPtr interface{}) {
	langTag := GetLangTag(context)
	if langTag == alc_lang.None {
		langTag = alc_lang.Cn
	}
	if bindErr := context.ShouldBindJSON(reqPtr); bindErr != nil {
		if vErrs, ok := bindErr.(validator.ValidationErrors); ok {
			for _, vErr := range vErrs {
				e := vErr.(validator.FieldError)
				var errMsg string
				if langTag == alc_lang.Cn {
					errMsg = e.Translate(alc_validator.TransCn)
				} else {
					errMsg = e.Translate(alc_validator.TransEn)
				}
				if filed, ok := reflect.TypeOf(reqPtr).Elem().FieldByName(e.StructField()); ok {
					var filedName string
					if langTag == alc_lang.Cn {
						filedName = filed.Tag.Get("cn")
					} else {
						filedName = filed.Tag.Get("en")
					}
					errMsg = strings.ReplaceAll(errMsg, e.Field(), filedName)
				}
				panic(alc_errs.CommonError(errMsg))
			}
		} else {
			panic(alc_errs.CommonError(bindErr.Error()))
		}
	}
}

func CheckDTO4FormMultipart(context *gin.Context, reqPtr interface{}) {
	langTag := GetLangTag(context)
	if langTag == alc_lang.None {
		langTag = alc_lang.Cn
	}
	if bindErr := context.ShouldBindWith(reqPtr, binding.FormMultipart); bindErr != nil {
		if vErrs, ok := bindErr.(validator.ValidationErrors); ok {
			for _, vErr := range vErrs {
				e := vErr.(validator.FieldError)
				var errMsg string
				if langTag == alc_lang.Cn {
					errMsg = e.Translate(alc_validator.TransCn)
				} else {
					errMsg = e.Translate(alc_validator.TransEn)
				}
				if filed, ok := reflect.TypeOf(reqPtr).Elem().FieldByName(e.StructField()); ok {
					var filedName string
					if langTag == alc_lang.Cn {
						filedName = filed.Tag.Get("cn")
					} else {
						filedName = filed.Tag.Get("en")
					}
					errMsg = strings.ReplaceAll(errMsg, e.Field(), filedName)
				}
				panic(alc_errs.CommonError(errMsg))
			}
		} else {
			panic(alc_errs.CommonError(bindErr.Error()))
		}
	}
}
