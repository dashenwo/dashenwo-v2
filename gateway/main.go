package main

import "github.com/dashenwo/dashenwo/v2/gateway/cmd"

func main() {
	if err := cmd.Init(); err != nil {
		panic("应用初始化失败" + err.Error())
	}
}
