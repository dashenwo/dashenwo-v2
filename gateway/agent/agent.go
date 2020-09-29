package agent

import (
	"github.com/dashenwo/dashenwo/v2/gateway/pkg/logger"
	"github.com/dashenwo/dashenwo/v2/gateway/schema"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

func Command() *cli.Command {
	command := &cli.Command{
		Name:  "agent",
		Usage: "Run the api gateway agent",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "address",
				Usage:   "Set the gateway agent address :10001",
				EnvVars: []string{"OKGATEWAY_ADDRESS"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return run(ctx)
		},
	}
	return command
}

func run(ctx *cli.Context) error {
	var state int32 = 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := Init(ctx)
	if err != nil {
		return err
	}

EXIT:
	for {
		sig := <-sc
		logger.Infof("接收到信号[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			atomic.CompareAndSwapInt32(&state, 1, 0)
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}
	cleanFunc()
	logger.Infof("服务退出")
	time.Sleep(time.Second)
	os.Exit(int(atomic.LoadInt32(&state)))
	return nil
}

// 初始化
func Init(ctx *cli.Context) (func(), error) {
	var err error
	//1.初始化配置文件

	//2.初始化etcd

	//3.同步配置

	//4.初始化http服务
	r := gin.Default()

	r.Any("/*path", func(c *gin.Context) {
		// init获取配置文件
		api := &schema.Api{
			Title:         "用户登录",
			Upstreams:     "192.168.3.4:8002",
			RequestMethod: []string{"POST", "GET"},
			RequestURL:    "/user/login",
			Proto:         "GRPC",
			TargetMethod:  "GRPC",
			TargetURL:     "com.dashenwo.srv.account_srv.Account/Login",
			TimeOut:       2000,
			Gid:           1,
			Pid:           1,
		}

		logger.Info("进来了")
	})
	if err = r.Run(":9000"); err != nil {

	}
	return func() {

	}, err
}
