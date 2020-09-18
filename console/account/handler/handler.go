package handler

import (
	"github.com/dashenwo/dashenwo/v2/console/account/internal/service"
	"github.com/dashenwo/dashenwo/v2/console/account/proto"
	"github.com/micro/go-micro/v2/server"
	"go.uber.org/dig"
)

func Apply(server server.Server, c *dig.Container) {
	_ = c.Invoke(func(accountService *service.AccountService) {
		_ = proto.RegisterAccountHandler(server, NewAccountHandler(accountService))
	})
}
