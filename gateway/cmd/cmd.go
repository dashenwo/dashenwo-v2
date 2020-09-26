package cmd

import (
	"github.com/dashenwo/dashenwo/v2/gateway/agent"
	"github.com/dashenwo/dashenwo/v2/gateway/server"
	"github.com/urfave/cli/v2"
	"os"
)

var VERSION = "1.0.0"

// 初始化方法
func Init() error {
	app := cli.NewApp()
	app.Name = "gateway"
	app.Usage = "api网关"
	app.Version = VERSION
	// 注册命令
	setupCommand(app)
	// 注册全局flags
	setupGlobalFlags(app)
	// 运行
	return app.Run(os.Args)
}

// 注册命令方法
func setupCommand(app *cli.App) {
	// 注册server
	app.Commands = append(app.Commands, server.Command())
	// 注册agent
	app.Commands = append(app.Commands, agent.Command())
}

// 注册全局flag
func setupGlobalFlags(app *cli.App) {
	// 设置默认的flag
	defaultFlags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "conf_path",
			Usage:   "配置文件，可以有多个，默认使用",
			EnvVars: []string{"OK_CONF_PATH"},
		},
		&cli.StringFlag{
			Name:    "address",
			Usage:   "服务端运行的地址",
			EnvVars: []string{"OK_ADDRESS"},
		},
		&cli.BoolFlag{
			Name:    "tls_enable",
			Usage:   "启用TLS支持。需要指定证书和密钥文件",
			EnvVars: []string{"OK_TLS_ENABLE"},
		},
		&cli.StringFlag{
			Name:    "tls_address",
			Usage:   "服务端启用tls时监听的地址",
			EnvVars: []string{"OK_TLS_ADDRESS"},
		},
		&cli.StringFlag{
			Name:    "tls_cert_file",
			Usage:   "TLS证书文件的路径",
			EnvVars: []string{"OK_TLS_CERT_FILE"},
		},
		&cli.StringFlag{
			Name:    "tls_key_file",
			Usage:   "TLS密钥文件的路径",
			EnvVars: []string{"OK_TLS_KEY_FILE"},
		},
	}
	// 把flag添加进app
	app.Flags = append(app.Flags, defaultFlags...)
}
