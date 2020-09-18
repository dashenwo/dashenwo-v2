package redis

import (
	"github.com/go-redis/redis"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/util/log"
	"sync"
)

type Redis struct {
	// redis地址
	Host string
	// 密码
	Password string
}

var (
	client *redis.Client
	conf   Redis
	once   sync.Once
	err    error
	pong   string
)

func Init() (*redis.Client, error) {
	once.Do(func() {
		conf = Redis{}
		err := config.Get("redis").Scan(&conf)
		client = redis.NewClient(&redis.Options{
			Addr:     conf.Host,
			Password: conf.Password, // no password set
		})
		pong, err = client.Ping().Result()
		log.Log("这是我打的日志", pong, err)
		if err != nil {
			log.Log("这是我打的日志", pong, err)
		}
	})
	return client, err
}
