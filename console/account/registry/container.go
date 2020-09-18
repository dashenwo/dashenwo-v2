package registry

import (
	"github.com/dashenwo/dashenwo/v2/console/account/global"
	"github.com/dashenwo/dashenwo/v2/console/account/internal/repository/persistence/gorm"
	"github.com/dashenwo/dashenwo/v2/console/account/internal/service"
	"github.com/dashenwo/dashenwo/v2/pkg/storage/elasticsearch"
	"github.com/dashenwo/dashenwo/v2/pkg/storage/redis"
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
	// 初始化redis
	if client, err := redis.Init(); err == nil {
		global.Redis = client
	} else {
		panic("初始化redis失败")
	}
	err2 := c.Provide(gorm.NewAccountRepository)
	log.Info(err2)
	err3 := c.Provide(service.NewAccountService)
	log.Info(err3)
}
