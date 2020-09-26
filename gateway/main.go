package gateway

import "github.com/micro/go-micro/v2/config/cmd"

func main() {
	if err := cmd.Init(); err != nil {
		panic("应用初始化失败" + err.Error())
	}
}
