package service

import (
	"context"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	conf "github.com/dashenwo/dashenwo/v2/console/captcha/config"
	"github.com/dashenwo/dashenwo/v2/console/captcha/internal/model"
	"github.com/dashenwo/dashenwo/v2/console/captcha/internal/repository"
	"github.com/dashenwo/dashenwo/v2/console/captcha/schema"
	"github.com/dashenwo/dashenwo/v2/console/snowflake/proto"
	"github.com/dashenwo/dashenwo/v2/pkg/utils/regexp"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/util/log"
	"gopkg.in/gomail.v2"
	"math/rand"
	"strconv"
	"time"
)

type CaptchaService struct {
	repo repository.CaptchaRepository
}

func NewCaptchaService(repo repository.CaptchaRepository) *CaptchaService {
	return &CaptchaService{
		repo: repo,
	}
}

// 生成验证码并发送
func (s CaptchaService) Generate(recipient string, recipientType int32) (*schema.Captcha, error) {
	var sendError error
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	rsp, err := getSnowflakeId()
	if err != nil {
		return nil, err
	}
	captcha := &model.Captcha{
		ID:        strconv.FormatInt(rsp.Id, 10),
		Code:      code,
		Recipient: recipient,
		Type:      int(recipientType),
	}
	if err := s.repo.Insert(captcha); err != nil {
		return nil, errors.New(conf.AppId, "获取验证码失败", 509)
	}
	// 判断账号类型
	if regexp.VerifyEmailFormat(recipient) {
		// 邮箱注册
		title := "欢迎使用酷答网"
		body := "欢迎使用酷答网，您的验证码为：" + code + "，如果非本人操作，请忽略"
		sendError = sendEmail(title, body, []string{recipient})
	} else if regexp.VerifyMobileFormat(recipient) {
		sendError = sendSms(recipient, code, recipientType)
	} else {
		return nil, errors.New(conf.AppId, "未识别到您的账号类型", 510)
	}
	if sendError != nil {
		return nil, errors.New(conf.AppId, sendError.Error(), 510)
	}
	return nil, nil
}

// 注册方法
func (s CaptchaService) Verify(id, recipient, code string) (*schema.Captcha, error) {

	return nil, nil
}

func getSnowflakeId() (*proto.Response, error) {
	now1 := time.Now().UnixNano() / 1e6
	service := micro.NewService()
	service.Init()
	log.Info("新建一个service获取到的地址", service.Client().Options().Registry.Options().Addrs)
	// create the proto client for helloworld
	srv := proto.NewSnowflakeService("com.dashenwo.srv.snowflake", service.Client())
	// call an endpoint on the service
	rsp, callErr := srv.Generate(context.Background(), &proto.Request{})
	if callErr != nil {
		return nil, errors.New("com.dashenwo.srv.snowflake", callErr.Error(), 506)
	}
	log.Info("使用proto方式调用消耗：", (time.Now().UnixNano()/1e6)-now1)
	return rsp, nil
}

// 发送短信
func sendSms(phone, code string, sendType int32) error {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI4Fyti1ETXnD2chtv95fy", "ozPuE0hyTex2TcorFpdgkJCouQXhCO")
	if err != nil {
		return err
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phone
	request.SignName = "酷答网"
	if sendType == 1 {
		request.TemplateCode = "SMS_203190166"
	} else if sendType == 2 {
		request.TemplateCode = "SMS_203190166"
	} else if sendType == 3 {
		request.TemplateCode = "SMS_203190161"
	} else if sendType == 4 {
		request.TemplateCode = "SMS_203190166"
	}
	request.TemplateParam = "{\"code\":\"" + code + "\"}"
	response, err1 := client.SendSms(request)
	if err1 != nil {
		return err1
	}
	if response.Code == "OK" {
		return nil
	}
	return nil
}

// 发送邮件
func sendEmail(title, body string, to []string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "admin@coolask.cn")
	m.SetHeader("To", to...)
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", title)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(conf.EmailConf.Host, conf.EmailConf.Port, conf.EmailConf.Username, conf.EmailConf.Password)
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
