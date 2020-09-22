package repository

import (
	"github.com/dashenwo/dashenwo/v2/console/captcha/internal/model"
)

// 用户接口
type CaptchaRepository interface {
	FindById(id string) (*model.Captcha, error)
	Insert(account *model.Captcha) error
}
