package handler

import (
	"context"
	"github.com/dashenwo/dashenwo/v2/console/snowflake/global"
	"github.com/dashenwo/dashenwo/v2/console/snowflake/proto"
	"github.com/micro/go-micro/v2/util/log"
)

type Snowflake struct {
}

func NewSnowflake() *Snowflake {
	return &Snowflake{}
}

func (h *Snowflake) Generate(ctx context.Context, req *proto.Request, res *proto.Response) error {
	log.Log("进来访问了")
	id := global.SnowflakeNode.Generate()
	res.Id = id.Int64()
	return nil
}
