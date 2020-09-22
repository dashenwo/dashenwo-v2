package handler

import (
	"github.com/dashenwo/dashenwo/v2/console/captcha/internal/service"
	"github.com/dashenwo/dashenwo/v2/console/captcha/proto"
	"github.com/micro/go-micro/v2/server"
	"go.uber.org/dig"
)

func Apply(server server.Server, c *dig.Container) {
	_ = c.Invoke(func(captchaService *service.CaptchaService) {
		_ = proto.RegisterCaptchaHandler(server, NewAccountHandler(captchaService))
	})
}
