package repository

import (
	"github.com/dashenwo/dashenwo/v2/console/captcha/internal/model"
)

// 用户接口
type CaptchaRepository interface {
	FindByModel(captcha *model.Captcha) (*model.Captcha, error)
	Insert(captcha *model.Captcha) error
}
