package handler

import (
	"context"
	"github.com/dashenwo/dashenwo/v2/console/snowflake/global"
	"github.com/dashenwo/dashenwo/v2/console/snowflake/proto"
)

type Snowflake struct {
}

func NewSnowflake() *Snowflake {
	return &Snowflake{}
}

func (h *Snowflake) Generate(ctx context.Context, req *proto.Request, res *proto.Response) error {
	id := global.SnowflakeNode.Generate()
	res.Id = id.Int64()
	return nil
}
