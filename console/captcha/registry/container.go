package registry

import (
	conf "github.com/dashenwo/dashenwo/v2/console/captcha/config"
	"github.com/dashenwo/dashenwo/v2/console/captcha/global"
	"github.com/dashenwo/dashenwo/v2/console/captcha/internal/repository/persistence/gorm"
	"github.com/dashenwo/dashenwo/v2/console/captcha/internal/service"
	"github.com/dashenwo/dashenwo/v2/pkg/storage/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/util/log"
	"go.uber.org/dig"
)

func NewContainer() (*dig.Container, error) {
	c := dig.New()
	buildCaptchaUsecase(c)
	return c, nil
}

func buildCaptchaUsecase(c *dig.Container) {
	// 初始化邮件配置信息
	if err := config.Get("email").Scan(&conf.EmailConf); err != nil {
		panic("初始化email配置信息出错")
	}
	// DB初始化
	gorm.InitDb()
	// 初始化redis
	if client, err := redis.Init(); err == nil {
		global.Redis = client
	} else {
		panic("初始化redis失败")
	}
	err2 := c.Provide(gorm.NewCaptchaRepository)
	log.Info(err2)
	err3 := c.Provide(service.NewCaptchaService)
	log.Info(err3)
}
