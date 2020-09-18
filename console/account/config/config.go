package config

import "time"

var (
	ConfPath = "./conf/"
	AppId    = "com.dashenwo.srv.account"
)

// 数据库配置信息
type Database struct {
	LogMode     bool
	AutoMigrate bool
	Engine      string
	Host        string
	Port        string
	User        string
	Password    string
	Name        string

	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}
