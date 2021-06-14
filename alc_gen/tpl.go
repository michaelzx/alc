package alc_gen

import (
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"strings"
	"text/template"
)

func loadTpl(tplStr string) (*template.Template, error) {
	t, err := template.New("tpl").Funcs(template.FuncMap{
		"ToLowerCamel": strcase.ToLowerCamel,
		"ToShort": func(s string) string {
			return strings.ToLower(string(s[0]))
		},
	}).Parse(tplStr)
	if err != nil {
		return nil, errors.Wrap(err, "模板引擎加载失败")
	}
	return t, nil
}
