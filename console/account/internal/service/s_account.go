package service

import (
	"context"
	conf "github.com/dashenwo/dashenwo/v2/console/account/config"
	"github.com/dashenwo/dashenwo/v2/console/account/internal/model"
	"github.com/dashenwo/dashenwo/v2/console/account/internal/repository"
	"github.com/dashenwo/dashenwo/v2/console/account/schema"
	"github.com/dashenwo/dashenwo/v2/console/captcha/proto"
	"github.com/dashenwo/dashenwo/v2/pkg/crypto"
	"github.com/dashenwo/dashenwo/v2/pkg/utils/generate"
	"github.com/jinzhu/copier"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/errors"
	"strconv"
	"time"
)

type AccountService struct {
	repo repository.UserRepository
}

func NewAccountService(repo repository.UserRepository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

// 登录方法
func (s AccountService) Login(username string, password string) (*schema.Account, error) {
	// 查询用户
	account, err := s.repo.FindByName(username)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, errors.New(conf.AppId, "账号或者密码错误", 501)
	}
	isLogin := crypto.ComparePasswords(account.Password, password+account.Salt)
	if isLogin == false {
		return nil, errors.New(conf.AppId, "账号或者密码错误", 502)
	}
	item := new(schema.Account)
	_ = copier.Copy(item, account)
	return item, err
}

// 注册方法
func (s AccountService) Register(nickname, password, recipient, code string) (*schema.Account, error) {
	//1.验证验证码是否正确
	service := micro.NewService()
	service.Init()
	srv := proto.NewCaptchaService("com.dashenwo.srv.captcha", service.Client())
	_, err := srv.Verify(context.Background(), &proto.VerifyRequest{Recipient: recipient, Code: code, Type: 2})
	if err != nil {
		return nil, err
	}
	//2.调用id
	rsp, callErr := generate.GetSnowflakeId()
	if callErr != nil {
		return nil, callErr
	}
	salt := crypto.GetRandomString(8)
	account := &model.Account{
		ID:           strconv.FormatInt(rsp.Id, 10),
		Nickname:     nickname,
		Phone:        recipient,
		Salt:         salt,
		Password:     crypto.HashAndSalt(password, salt),
		RegisterTime: time.Now().Unix(),
	}
	err = s.repo.Insert(account)
	if err != nil {
		return nil, err
	}
	item := new(schema.Account)
	_ = copier.Copy(item, account)
	return item, nil
}
