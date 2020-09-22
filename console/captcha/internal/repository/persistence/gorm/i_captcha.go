package gorm

import (
	"github.com/dashenwo/dashenwo/v2/console/captcha/config"
	"github.com/dashenwo/dashenwo/v2/console/captcha/internal/model"
	"github.com/dashenwo/dashenwo/v2/console/captcha/internal/repository"
	"github.com/micro/go-micro/v2/errors"
)

type CaptchaRepository struct {
}

func NewCaptchaRepository() repository.CaptchaRepository {
	return &CaptchaRepository{}
}

func (a *CaptchaRepository) FindById(id string) (*model.Captcha, error) {
	account := model.Captcha{}
	if err := db.Where("id = ?", id).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *CaptchaRepository) Insert(account *model.Captcha) error {
	if err := db.Create(account).Error; err != nil {
		return errors.New(config.AppId, err.Error(), 201)
	}
	return nil
}
