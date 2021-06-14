package internal
{{- $length := len .Imports -}}
{{if ne $length 0}}
import ({{range .Imports}}
    {{.}}{{end}}
){{end}}

type {{.StructName}} struct { {{range .Fields}}
    {{.}}{{end}}
}

func ({{.StructName}}) TableName() string {
    return "{{.TableName}}"
}