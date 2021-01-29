package alc_config

import (
	"github.com/spf13/viper"
	"path/filepath"
)

type MixinConfig interface {
	// 加载配置，用什么实现可自行定义。下面提供了加载toml的加载方式
	Load(configFilePath string)
}

func LoadYaml(path string, cfg MixinConfig) error {
	// 读取yaml文件
	v := viper.New()
	// 设置读取的配置文件名
	v.SetConfigName(filepath.Base(path))
	// windows环境下为%GOPATH，linux环境下为$GOPATH
	v.AddConfigPath(filepath.Dir(path))
	// 设置配置文件类型
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	// 也可以直接反序列化为Struct
	if err := v.Unmarshal(cfg); err != nil {
		return err
	}
	return nil
}
