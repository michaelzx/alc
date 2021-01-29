package configs

import (
	"alchemy/alc/alc_config"
)

type Config struct {
	Mysql alc_config.MysqlConfig
}

func (m *Config) Load(configFilePath string) {
	err := alc_config.LoadYaml(configFilePath, m)
	if err != nil {
		panic(err)
	}
}
