package global

import (
	"github.com/dashenwo/go-library/session"
	"github.com/go-redis/redis/v8"
	"github.com/micro/go-micro/v2/client"
	"github.com/olivere/elastic/v7"
)

var (
	// elasticsearch客户端
	Es            *elastic.Client
	Redis         redis.UniversalClient
	SessionManage *session.Manager
	RequestClient client.Client
)
