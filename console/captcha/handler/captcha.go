package handler

import (
	"context"
	conf "github.com/dashenwo/dashenwo/v2/console/captcha/config"
	"github.com/dashenwo/dashenwo/v2/console/captcha/global"
	"github.com/dashenwo/dashenwo/v2/console/captcha/internal/service"
	"github.com/dashenwo/dashenwo/v2/console/captcha/proto"
	"github.com/dashenwo/dashenwo/v2/pkg/utils/validate"
	"github.com/micro/go-micro/v2/errors"
	"time"
)

type Captcha struct {
	captchaService *service.CaptchaService
}

// 实例化方法
func NewAccountHandler(captcha *service.CaptchaService) *Captcha {
	return &Captcha{
		captchaService: captcha,
	}
}

// 登录handler
func (a *Captcha) Generate(ctx context.Context, req *proto.GenerateRequest, res *proto.GenerateResponse) error {
	//1.验证数据
	if err := validate.Validate(req, conf.AppId); err != nil {
		return err
	}
	captcha, err := a.captchaService.Generate(req.Recipient, req.Type)
	if err != nil {
		return err
	}
	if err := global.Redis.Set(captcha.ID, captcha.Code, time.Minute*10).Err(); err != nil {
		return errors.New(conf.AppId, "验证码存储失败", 506)
	}
	now := time.Now()
	hh, _ := time.ParseDuration("10m")
	res.Id = captcha.ID
	res.Expires = now.Add(hh).Format("2006-01-02 15:04:05")
	return nil
}

// 注册handler
func (a *Captcha) Verify(ctx context.Context, req *proto.VerifyRequest, res *proto.VerifyResponse) error {

	return nil
}
