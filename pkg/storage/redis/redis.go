package redis

import (
	"github.com/go-redis/redis/v8"
	"sync"
)

type Redis struct {
	// redis地址
	Hosts []string
	// 密码
	Password string
}

var (
	client redis.UniversalClient
	conf   Redis
	once   sync.Once
	err    error
)

func Init() (redis.UniversalClient, error) {
	once.Do(func() {

	})
	return client, err
}
