package global

import (
	"github.com/go-redis/redis"
	"github.com/micro/go-micro/v2/client"
	"github.com/olivere/elastic/v7"
)

var (
	// elasticsearch客户端
	Es            *elastic.Client
	Redis         *redis.Client
	RequestClient client.Client
)
