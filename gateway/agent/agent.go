package agent

import (
	"github.com/micro/go-micro/v2/logger"
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
	//1.初始化配置文件

	//2.初始化etcd

	//3.同步配置
}
