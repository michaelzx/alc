package alc_validator

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"regexp"
	"sync"
)

var TransCn ut.Translator
var TransEn ut.Translator

type DefaultValidator struct {
	uni            *ut.UniversalTranslator
	once           sync.Once
	validate       *validator.Validate
	transList      []func()
	validationList []func()
}

// ValidateStruct receives any kind of type, but only performed struct or pointer to struct type.
func (v *DefaultValidator) ValidateStruct(obj interface{}) error {
	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	if valueType == reflect.Struct {
		v.lazyInit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

// Engine returns the underlying validator engine which powers the default
// Validator instance. This is useful if you want to register custom validations
// or struct level validations. See validator GoDoc for more info -
// https://godoc.org/gopkg.in/go-playground/validator.v8
func (v *DefaultValidator) Engine() interface{} {
	v.lazyInit()
	return v.validate
}

var (
	uni *ut.UniversalTranslator
)

func (v *DefaultValidator) TransCn(tag string, registerTranslationsFunc validator.RegisterTranslationsFunc, translationFunc validator.TranslationFunc) {
	v.transList = append(v.transList, func() {
		_ = v.validate.RegisterTranslation(tag, TransCn, registerTranslationsFunc, translationFunc)
	})
}
func (v *DefaultValidator) TransEn(tag string, registerTranslationsFunc validator.RegisterTranslationsFunc, translationFunc validator.TranslationFunc) {
	v.transList = append(v.transList, func() {
		_ = v.validate.RegisterTranslation(tag, TransEn, registerTranslationsFunc, translationFunc)
	})
}
func (v *DefaultValidator) RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) {
	v.validationList = append(v.validationList, func() {
		_ = v.validate.RegisterValidation(tag, fn, callValidationEvenIfNull...)
	})
}
func (v *DefaultValidator) lazyInit() {
	v.once.Do(func() {

		uni = ut.New(en.New(), zh.New())

		// this is usually know or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)
		TransCn, _ = uni.GetTranslator("zh")
		TransEn, _ = uni.GetTranslator("en")

		v.validate = validator.New()
		_ = zhTranslations.RegisterDefaultTranslations(v.validate, TransCn)
		_ = enTranslations.RegisterDefaultTranslations(v.validate, TransEn)

		_ = v.validate.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
			v := fl.Field().String()
			reg := `^1\d{10}$`
			rgx := regexp.MustCompile(reg)
			return rgx.MatchString(v)
		})
		_ = v.validate.RegisterValidation("admin_account", func(fl validator.FieldLevel) bool {
			v := fl.Field().String()
			reg := `^[a-zA-Z0-9_-]{4,16}$`
			rgx := regexp.MustCompile(reg)
			return rgx.MatchString(v)
		})
		_ = v.validate.RegisterValidation("admin_password", func(fl validator.FieldLevel) bool {
			v := fl.Field().String()
			if v == "" {
				return true
			}
			rgx := regexp.MustCompile(`[A-Za-z]`)
			if !rgx.MatchString(v) {
				return false
			}
			rgx = regexp.MustCompile(`[0-9]`)
			if !rgx.MatchString(v) {
				return false
			}
			if len(v) < 6 || len(v) > 20 {
				return false
			}
			return true
		})
		// 可以用以下方式来增加验证器自定义规则
		for i, _ := range v.validationList {
			v.validationList[i]()
		}
		// 可以用以下方式来增加验证器自定义规则的翻译
		for i, _ := range v.transList {
			v.transList[i]()
		}
		_ = v.validate.RegisterTranslation("admin_password", TransCn, func(ut ut.Translator) error {
			return ut.Add("admin_password", "{0} 必须包含数字和英文，长度在6~20!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("admin_password", fe.Field())
			return t
		})
		// _ = v.validate.RegisterValidation("dir_name", func(fl validator.FieldLevel) bool {
		// 	v := fl.Field().String()
		// 	if v == "" {
		// 		return true
		// 	}
		// 	rgx := regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
		// 	if !rgx.MatchString(v) {
		// 		return false
		// 	}
		// 	return true
		// })
		// // 可以用以下方式来增加验证器自定义规则的翻译
		// _ = v.validate.RegisterTranslation("dir_name", TransCn, func(ut ut.Translator) error {
		// 	return ut.Add("dir_name", "{0} 只能包含a-z,A-Z,0-9,-,_", true) // see universal-translator for details
		// }, func(ut ut.Translator, fe validator.FieldError) string {
		// 	t, _ := ut.T("dir_name", fe.Field())
		// 	return t
		// })
		_ = v.validate.RegisterValidation("int", func(fl validator.FieldLevel) bool {
			p := fl.Param()
			v := fl.Field().String()
			reg := `^\d+$`
			if p != "" {
				reg = fmt.Sprintf(`^\d{%s}$`, p)
			}
			rgx := regexp.MustCompile(reg)
			return rgx.MatchString(v)
		})
		v.validate.SetTagName("valid")
	})
}
