package alc_gin

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/michaelzx/alc/alc_errs"
	"github.com/michaelzx/alc/alc_lang"
	"github.com/michaelzx/alc/alc_validator"
	"reflect"
	"strings"
)

func CheckStruct(reqPtr interface{}, langTags ...alc_lang.Tag) error {
	if len(langTags) > 1 {
		panic(alc_errs.CommonError("langTags参数数量，仅限：0或1个"))
	}
	langTag := alc_lang.Cn
	if len(langTags) == 1 {
		langTag = langTags[0]
	}

	if bindErr := binding.Validator.ValidateStruct(reqPtr); bindErr != nil {
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
				return errors.New(errMsg)
			}
		} else {
			return bindErr
		}
	}
	return nil
}
