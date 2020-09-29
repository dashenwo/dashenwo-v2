package agent

import (
	"github.com/dashenwo/dashenwo/v2/gateway/pkg/logger"
	_ "github.com/dashenwo/dashenwo/v2/gateway/schema"
	"github.com/urfave/cli/v2"
	//"github.com/valyala/fasthttp"
	"net/http"
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
	//1.初始化配置文件,注册路由

	//2.初始化etcd

	//3.同步配置

	//4.初始化http服务
	http.HandleFunc("/", func(witer http.ResponseWriter, request *http.Request) {
		witer.Write([]byte("这就是了"))
	})
	_ = http.ListenAndServe(":9001", nil)
	//err = fasthttp.ListenAndServe(":9000", func(ctx *fasthttp.RequestCtx) {
	//	_, _ = ctx.WriteString("这就是了")
	//})
	return func() {

	}, err
}
