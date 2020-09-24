package config

import (
	"github.com/dashenwo/dashenwo/v2/pkg/storage/kafka"
	"time"
)

var (
	ConfPath  = "./conf/"
	AppId     = "com.dashenwo.srv.captcha"
	EmailConf = Email{}
	KafkaConf = kafka.Kafka{}
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

type Email struct {
	Host     string
	Port     int
	Username string
	Password string
}
