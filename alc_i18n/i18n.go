package alc_i18n

import (
	"bytes"
	"path/filepath"
	"strings"
	"text/template"
)

const errPrefix = "err: "

var localesDir string

type CommonMap = map[string]interface{}

func Init(dir string) {
	localesDir = dir
	initTransMapCache()
}

type I18n struct {
	keyPath string
	locale  Locale
	tag     string
}

func (i18n *I18n) Tag(tag string) *I18n {
	i18n.tag = tag
	return i18n
}

func (i18n *I18n) parseFilePath() string {
	filePath := filepath.Join(localesDir, i18n.parseFileName())
	return filePath
}
func (i18n *I18n) parseFileName() string {
	fileName := i18n.locale
	if i18n.tag != "" {
		fileName += "." + i18n.tag
	}
	fileName += ".json"
	return fileName
}
func (i18n *I18n) getTransMap() CommonMap {
	filePath := i18n.parseFilePath()
	if v, exists := transMapCache[filePath]; exists {
		return v
	}
	return nil
}

func (i18n *I18n) Trans() string {
	// 先从缓存中取
	cacheKey := i18n.parseFileName() + "->" + i18n.keyPath
	localeString := getTransKeyCache(cacheKey)
	if localeString != "" {
		return localeString
	}
	// 解析成map
	jsonMap := i18n.getTransMap()
	if jsonMap == nil {
		return errPrefix + "[" + i18n.parseFileName() + "]" + " does not exist"
	}
	// 根据路径，获取具体的文字
	keys := strings.Split(i18n.keyPath, ".")
	var currentValue interface{}
	currentValue = jsonMap
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		if nextMap, ok := currentValue.(CommonMap); ok {
			if nextValue, exists := nextMap[key]; exists {
				currentValue = nextValue
			}
		} else {
			return errPrefix + "[" + strings.Join(keys[:i], ".") + "]" + " is not object"
		}

	}
	// 最后的值，必须是字符串
	if localeString, ok := currentValue.(string); ok {
		addTransKeyCache(cacheKey, localeString)
		return localeString
	} else {
		return errPrefix + "[" + i18n.keyPath + "]" + " is not string"
	}
}

func (i18n *I18n) TransWithValues(values interface{}) string {
	localeString := i18n.Trans()
	tmplCacheKey := i18n.parseFileName() + "->" + i18n.keyPath
	tmpl := getTmplCache(tmplCacheKey)
	if tmpl == nil {
		var err error
		tmpl, err = template.New(i18n.parseFileName() + "->" + i18n.keyPath).Option("missingkey=error").Parse(localeString)
		if err != nil {
			return errPrefix + err.Error()
		}
		addTmplCache(tmplCacheKey, tmpl)
	}
	var b bytes.Buffer
	err := tmpl.Execute(&b, values)
	if err != nil {
		return errPrefix + err.Error()
	}
	return string(b.Bytes())
}
