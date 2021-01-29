package alc_lang

import (
	"strings"
)

const GinContextKey = "LangTag"

type Tag string

func (t Tag) String() string {
	return string(t)
}
func TagFromString(s string) Tag {
	lowerV := strings.ToLower(s)
	for _, langTag := range Tags {
		if lowerV == langTag.String() {
			return langTag
		}
	}
	return None
}

const (
	None Tag = ""
	Cn   Tag = "cn"
	En   Tag = "en"
)

var Tags = []Tag{
	Cn,
	En,
}
