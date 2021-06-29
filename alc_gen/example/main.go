package main

import (
	"github.com/michaelzx/alc/alc_gen"
	"github.com/michaelzx/alc/alc_gen/example/config"
	"log"
	"path/filepath"
)

func main() {
	cfg := new(config.Config)
	cfg.Load("./config/config.yaml")
	rootPath, err := filepath.Abs("./")
	if err != nil {
		log.Fatal(err)
	}
	g, err := alc_gen.New(alc_gen.Config{
		RootPath:    rootPath,
		RootPackage: "github.com/michaelzx/alc/alc_gen/example",
		DbCfg:       cfg.Mysql,
		Tables: []string{
			// "member",
			// "team",
			// "meta",
			// "dict",
			// "module",
			"module_data",
			"module_model",
			"module_model_mgr",
			"module_node",
			"module_allow",
			"module_group",
			"module_group_mapping",
			"module_group_mapping",
		},
		TablePrefix: "module_",
		GenModel:    true,
	})
	if err != nil {
		log.Fatal(err)
	}
	g.Run()
}
