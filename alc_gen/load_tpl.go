package alc_gen

import (
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func (g *Generator) loadTpl(tplPath string) *template.Template {
	f, err := os.Open(tplPath)
	if err != nil {
		panic(errors.Wrap(err, "模板文件不存在"))
	}
	defer f.Close()

	var fd []byte
	fd, err = ioutil.ReadAll(f)
	if err != nil {
		panic(errors.Wrap(err, "读取模板文件失败"))
	}
	tplStr := string(fd)
	t, err := template.New("entTpl").Funcs(template.FuncMap{
		"ToLowerCamel": strcase.ToLowerCamel,
		"ToShort": func(s string) string {
			return strings.ToLower(string(s[0]))
		},
	}).Parse(tplStr)
	if err != nil {
		panic(errors.Wrap(err, "模板引擎加载失败"))
	}
	return t
}
