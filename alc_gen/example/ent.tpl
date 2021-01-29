package {{.PackageName}}
{{- $length := len .Imports -}}
{{if ne $length 0}}
import ({{range .Imports}}
    {{.}}{{end}}
){{end}}

type {{.StructName}}Ent struct { {{range .Fields}}
    {{.}}{{end}}
}

func ({{.StructName}}Ent) TableName() string {
    return "{{.TableName}}"
}