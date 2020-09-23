package generate

import (
	"context"
	"github.com/dashenwo/dashenwo/v2/console/snowflake/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/errors"
)

func GetSnowflakeId() (*proto.Response, error) {
	service := micro.NewService()
	service.Init()
	// create the proto client for helloworld
	srv := proto.NewSnowflakeService("com.dashenwo.srv.snowflake", service.Client())
	// call an endpoint on the service
	rsp, callErr := srv.Generate(context.Background(), &proto.Request{})
	if callErr != nil {
		return nil, errors.New("com.dashenwo.srv.snowflake", callErr.Error(), 506)
	}
	return rsp, nil
}
