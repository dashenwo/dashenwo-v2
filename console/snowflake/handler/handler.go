package handler

import (
	"github.com/dashenwo/dashenwo/v2/console/snowflake/proto"
	"github.com/micro/go-micro/v2/server"
	"go.uber.org/dig"
)

func Apply(server server.Server, c *dig.Container) {
	_ = c.Invoke(func() {
		_ = proto.RegisterSnowflakeHandler(server, NewSnowflake())
	})
}
