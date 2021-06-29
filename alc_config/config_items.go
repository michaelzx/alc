package alc_config

import "time"

// WebServerConfig Web服务器配置
type WebServerConfig struct {
	Port            int
	BodyLimit       int // 单位：M
	ShutDownWaitSec int
}

// AppConfig 应用全局配置
type AppConfig struct {
	Name    string // 程序中写死
	Version string // 程序中写死
	BaseUrl string
	CdnUrl  string
	Theme   string
}

// LoggerConfig 日志配置
type LoggerConfig struct {
	Mode string // prod or dev
}

// SnowflakeConfig Snowflake配置
type SnowflakeConfig struct {
	Node int64
}

// MysqlConfig MySQL配置
type MysqlConfig struct {
	Host          string
	Port          int
	DbName        string
	TablePrefix   string
	Usr           string
	Psw           string
	MaxConnection int
	MaxIdleConns  int
	MaxLifetime   time.Duration
	MaxOpenConns  int
	Debug         bool
}

// EmailConfig Email配置
type EmailConfig struct {
	From  string
	Title string
	Host  string
	Port  int
	Usr   string
	Psw   string
}
