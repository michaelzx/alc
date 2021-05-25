package alc_i18n

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"
)

// **************************************************************************************************
// tmplCache
// **************************************************************************************************
var tmplCache = sync.Map{}

func getTmplCache(key string) *template.Template {
	if cacheV, exists := tmplCache.Load(key); exists {
		if v, ok := cacheV.(*template.Template); ok {
			return v
		}
	}
	return nil
}

func addTmplCache(key string, tmpl *template.Template) {
	tmplCache.Store(key, tmpl)
}

// **************************************************************************************************
// transKeyCache
// **************************************************************************************************
var transKeyCache = sync.Map{}

func getTransKeyCache(key string) string {
	if cacheV, exists := transKeyCache.Load(key); exists {
		if v, ok := cacheV.(string); ok {
			return v
		}
	}
	return ""
}
func addTransKeyCache(key string, localeString string) {
	transKeyCache.Store(key, localeString)
}
func CleanTransKeyCache() {
	transKeyCache = sync.Map{}
}

// **************************************************************************************************
// transMapCache
// **************************************************************************************************
var transMapCache = make(map[string]CommonMap)

// initTransMapCache 初始化的时候，将所有json文件进行解析，并载入内存，避免每次都去读文件
func initTransMapCache() {
	files, err := os.ReadDir(localesDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		jsonMap := make(CommonMap)
		filePath := filepath.Join(localesDir, f.Name())
		if strings.HasSuffix(filePath, ".json") {
			fileBytes, err := ioutil.ReadFile(filePath)
			err = json.Unmarshal(fileBytes, &jsonMap)
			if err != nil {
				log.Fatal(err)
			}
			transMapCache[filePath] = jsonMap
		}
	}
}
