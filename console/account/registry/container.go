package registry

import (
	"github.com/dashenwo/dashenwo/v2/console/account/global"
	"github.com/dashenwo/dashenwo/v2/console/account/internal/repository/persistence/gorm"
	"github.com/dashenwo/dashenwo/v2/console/account/internal/service"
	"github.com/dashenwo/dashenwo/v2/pkg/storage/elasticsearch"
	"github.com/dashenwo/go-library/session"
	"github.com/dashenwo/go-library/session/storage"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro/v2/util/log"
	"go.uber.org/dig"
)

func NewContainer() (*dig.Container, error) {
	c := dig.New()
	buildAccountUsecase(c)
	return c, nil
}

func buildAccountUsecase(c *dig.Container) {
	// DB初始化
	gorm.InitDb()
	// 初始化elasticsearch
	if es, err := elasticsearch.Init(); err == nil {
		global.Es = es
	} else {
		panic("初始化es失败")
	}
	// 初始化session
	// 获取配置信息
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{"192.168.3.4:7001", "192.168.3.4:7002", "192.168.3.4:7003", "192.168.3.4:7004", "192.168.3.4:7005", "192.168.3.4:7006"},
		Password: "Liuqin76624291",
	})
	global.SessionManage = session.NewSessionManage(
		session.Storage(
			storage.NewRedisStorage(rdb),
		),
	)
	err2 := c.Provide(gorm.NewAccountRepository)
	log.Info(err2)
	err3 := c.Provide(service.NewAccountService)
	log.Info(err3)
}
