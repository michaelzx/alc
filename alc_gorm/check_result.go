package alc_gorm

import (
	"errors"
	"gorm.io/gorm"
)

// CheckResult 检查result
// 1.有记录被找到; 2.没错误,如果有错误直接panic
func CheckResult(result *gorm.DB) bool {
	switch {
	case errors.Is(result.Error, gorm.ErrRecordNotFound):
		return false
	case result.Error != nil:
		panic(result.Error)
	default:
		return true
	}
}

func HasUnknownError(result *gorm.DB) bool {
	return result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound)
}
