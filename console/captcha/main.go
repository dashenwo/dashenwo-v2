package main

import (
	conf "github.com/dashenwo/dashenwo/v2/console/captcha/config"
	"github.com/dashenwo/dashenwo/v2/console/captcha/handler"
	"github.com/dashenwo/dashenwo/v2/console/captcha/registry"
	tracer "github.com/dashenwo/dashenwo/v2/pkg/opentracing"
	"github.com/dashenwo/plugins/logger/zap/v2"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	zap2 "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	// 初始化日志库
	encodingConfig := zap2.NewProductionEncoderConfig()
	// 时间格式化
	encodingConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	l, err := zap.NewLogger(
		zap.WithCallerSkip(4),
		zap.WithConfig(zap2.NewProductionConfig()),
		zap.WithEncoderConfig(encodingConfig),
	)
	if err != nil {
		log.Fatal(err)
	}
	logger.DefaultLogger = l

	md := make(map[string]string)
	md["chain"] = "gray"

	//创建服务
	service := micro.NewService(
		micro.Name(conf.AppId),
		micro.Version("latest"),
		micro.Metadata(md),
		// 设置启动ip
		//micro.Address(":8002"),
		micro.Flags(
			&cli.StringFlag{
				Name:  "conf_path",
				Value: "./conf/",
				Usage: "配置文件目录",
			},
		),
		micro.Action(func(ctx *cli.Context) error {
			confPath := ctx.String("conf_path")
			conf.ConfPath = confPath
			// 配置加载
			err := config.LoadFile(conf.ConfPath + "config.yaml")
			return err
		}),
	)

	// 链路追踪
	t, closer, err := tracer.NewJaegerTracer(conf.AppId, "203.195.200.40:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()
	service.Init(
		// Tracing仅由Gateway控制，在下游服务中仅在有Tracing时启动
		micro.WrapCall(opentracing.NewCallWrapper(t)),
		micro.WrapHandler(opentracing.NewHandlerWrapper(t)),
	)
	// 初始化服务
	service.Init()

	c, err := registry.NewContainer()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}

	// Register Handler
	handler.Apply(service.Server(), c)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
