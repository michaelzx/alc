package {{.PackageName}}

var {{.StructName|ToLowerCamel}}Singleton *{{.StructName|ToLowerCamel}}Service
func init() {
    {{.StructName|ToLowerCamel}}Singleton = &{{.StructName|ToLowerCamel}}Service{}
}

type {{.StructName|ToLowerCamel}}Service struct {
}
func {{.StructName}}Service() *{{.StructName|ToLowerCamel}}Service {
    return {{.StructName|ToLowerCamel}}Singleton
}