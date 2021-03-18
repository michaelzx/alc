package configs

import "github.com/michaelzx/alc/alc_config"

var Mysql alc_config.MysqlConfig

func Init(configFilePath string) {
	cfg := &Config{}
	cfg.Load(configFilePath)
	Mysql = cfg.Mysql
}
