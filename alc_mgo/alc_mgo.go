package alc_mgo

import (
	"github.com/michaelzx/alc/alc_gorm"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

func Page(query qmgo.QueryI, page alc_gorm.PageParams, list interface{}) (pageVo *alc_gorm.PageVO, err error) {
	total, err := query.Count()
	if err != nil {
		return
	}
	pagination := alc_gorm.NewPagination(page.GetPageNum(), page.GetPageSize())
	pagination.Compute(total)

	err = query.Skip(pagination.GetSkipRows()).Limit(pagination.PageSize).All(&list)
	if err != nil {
		return
	}
	pageVo = &alc_gorm.PageVO{
		Pagination: pagination,
		List:       list,
	}
	return
}
func PageA()  {

}
func DRegex(pattern, option string) bson.D {
	return bson.D{{"$regex", primitive.Regex{Pattern: pattern, Options: option}}}
}

func DLike(words string) bson.D {
	return DRegex(ReplaceRegexSymbol(words), "ig")
}

func ReplaceRegexSymbol(regexPattern string) string {
	regexPattern = strings.ReplaceAll(regexPattern, `\`, `\\`)
	regexPattern = strings.ReplaceAll(regexPattern, `^`, `\^`)
	regexPattern = strings.ReplaceAll(regexPattern, `$`, `\$`)
	regexPattern = strings.ReplaceAll(regexPattern, `.`, `\.`)
	regexPattern = strings.ReplaceAll(regexPattern, `*`, `\*`)
	regexPattern = strings.ReplaceAll(regexPattern, `+`, `\+`)
	regexPattern = strings.ReplaceAll(regexPattern, `?`, `\?`)
	regexPattern = strings.ReplaceAll(regexPattern, `{`, `\{`)
	regexPattern = strings.ReplaceAll(regexPattern, `}`, `\}`)
	regexPattern = strings.ReplaceAll(regexPattern, `(`, `\(`)
	regexPattern = strings.ReplaceAll(regexPattern, `)`, `\)`)
	regexPattern = strings.ReplaceAll(regexPattern, `[`, `\[`)
	regexPattern = strings.ReplaceAll(regexPattern, `]`, `\]`)
	regexPattern = strings.ReplaceAll(regexPattern, `|`, `\|`)
	return regexPattern
}

func HasUnknownError(err error) bool {
	return err != nil && err != mongo.ErrNoDocuments
}
