package registry

import (
	"github.com/bwmarrin/snowflake"
	"github.com/dashenwo/dashenwo/v2/console/snowflake/config"
	"github.com/dashenwo/dashenwo/v2/console/snowflake/global"
	"github.com/dashenwo/dashenwo/v2/pkg/utils/iputils"
	"github.com/micro/go-micro/v2/server"
	"go.uber.org/dig"
	"strconv"
	"strings"
)

func NewContainer(server server.Server) (*dig.Container, error) {
	c := dig.New()
	buildSnowflakeUsecase(c, server)
	return c, nil
}

func buildSnowflakeUsecase(c *dig.Container, server server.Server) {
	// 定义一个nodeIndex
	nodeIndex := 1
	localIps := strings.Split(server.Options().Address, ":")
	services, err := server.Options().Registry.GetService(config.AppId)
	// 如果出错，当没有获取到services的时候则表示刚启动了第一台
	if err != nil {
		// 如果没有当前服务
		if err.Error() == "service not found" {
			nodeIndex = 1
		}
	} else {
		for _, service := range services {
			for _, node := range service.Nodes {
				// 把当前node的ip和端口分开
				ips := strings.Split(node.Address, ":")
				// 获取当前node的metadata信息
				metadata := node.Metadata
				// 获取当前node的nodeNo
				tempIndex, _ := strconv.Atoi(metadata["nodeNo"])
				// 如果是当前服务重启，并且还没有销毁就使用原来的nodeIndex
				if isLoacl, err := iputils.IsLocalIp(ips[0]); err == nil && isLoacl && localIps[1] == ips[1] {
					nodeIndex = tempIndex
					break
				} else {
					// 如果当前的nodeIndex被正在使用
					if nodeIndex == tempIndex {
						nodeIndex = nodeIndex + 1
					}
				}
			}
		}
	}
	// 设置metadata
	server.Options().Metadata["nodeNo"] = strconv.Itoa(nodeIndex)
	// 初始化雪花id
	if node, err := snowflake.NewNode(int64(nodeIndex)); err != nil {
		panic("初始化雪花node失败，错误信息为" + err.Error())
	} else {
		global.SnowflakeNode = node
	}
}
