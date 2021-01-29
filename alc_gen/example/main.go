package main

import (
	"alchemy/alc/alc_gen"
	"alchemy/alc/alc_gen/example/configs"
)

func main() {
	genApiV1()
}

func genApiV1() {
	configs.Init("configs/config.yaml")

	g := alc_gen.Generator{
		DbCfg:       configs.Mysql,
		BasePackage: "gitee.com/zx-io/paladin-go",
		TypePackage: "zky4/pld/alc_types",
		Tables: []string{
			// "admin",
			// "team",
			// "meta",
			// "dict",
		},
		OutDir:         "internal/app",
		PackageEnt:     "ent",
		PackageService: "service",
		TplEnt:         "cmd/gen-ent/ent.tpl",
		TplService:     "cmd/gen-ent/service.tpl",
		GenEnt:         true,
		GenService:     true,
		StructNameProcessor: func(structName string) string {
			// structName = strings.TrimPrefix(structName, "Album")
			return structName
		},
		FileNameProcessor: func(fileName string) string {
			// fileName = strings.TrimPrefix(fileName, "album_")
			return fileName
		},
	}
	g.Gen()
}
