package service

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	conf "github.com/dashenwo/dashenwo/v2/console/captcha/config"
	"github.com/dashenwo/dashenwo/v2/console/captcha/global"
	"github.com/dashenwo/dashenwo/v2/console/captcha/internal/model"
	"github.com/dashenwo/dashenwo/v2/console/captcha/internal/repository"
	"github.com/dashenwo/dashenwo/v2/console/captcha/schema"
	"github.com/dashenwo/dashenwo/v2/pkg/utils/generate"
	"github.com/dashenwo/dashenwo/v2/pkg/utils/regexp"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
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
	rsp, err := generate.GetSnowflakeId()
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

	// 保存到数据库后发送到kafka
	jsonByte, _ := json.Marshal(captcha)
	if err := sendMessage(conf.KafkaConf.Topic, string(jsonByte)); err != nil {
		return nil, errors.New(conf.AppId, "获取验证码失败", 509)
	}

	//// 判断账号类型
	//if regexp.VerifyEmailFormat(recipient) {
	//	// 邮箱注册
	//	title := "欢迎使用酷答网"
	//	body := "欢迎使用酷答网，您的验证码为：" + code + "，如果非本人操作，请忽略"
	//	sendError = sendEmail(title, body, []string{recipient})
	//} else if regexp.VerifyMobileFormat(recipient) {
	//	sendError = sendSms(recipient, code, recipientType)
	//} else {
	//	return nil, errors.New(conf.AppId, "未识别到您的账号类型", 510)
	//}
	//if sendError != nil {
	//	return nil, errors.New(conf.AppId, sendError.Error(), 510)
	//}
	item := new(schema.Captcha)
	_ = copier.Copy(item, captcha)
	return item, nil
}

// 注册方法
func (s CaptchaService) Verify(recipient, code string, recipientType int32) (*schema.Captcha, error) {

	captcha, err := s.repo.FindByModel(&model.Captcha{Recipient: recipient, Code: code, Type: int(recipientType)})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New(conf.AppId, "验证码错误", 511)
		} else {
			return nil, errors.New(conf.AppId, err.Error(), 511)
		}
	}
	if captcha == nil {
		return nil, errors.New(conf.AppId, "验证码错误", 512)
	}
	if captcha.CreatedTime+60*30 < int32(time.Now().Unix()) {
		return nil, errors.New(conf.AppId, "验证码已过期，请重新获取", 512)
	}
	item := new(schema.Captcha)
	_ = copier.Copy(item, captcha)
	return item, nil
}

func sendMessage(topic, msg string) error {
	//构建发送的消息，
	msg1 := &sarama.ProducerMessage{
		Topic: topic, //包含了消息的主题
		Value: sarama.StringEncoder(msg),
	}
	_, _, err := global.Kafka.SendMessage(msg1)
	if err != nil {
		log.Info("消息发送失败", err)
		return err
	}
	return nil
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
	if sendType == 1 { //登录
		request.TemplateCode = "SMS_203196967"
	} else if sendType == 2 { //注册
		request.TemplateCode = "SMS_203180176"
	} else if sendType == 3 { //修改密码
		request.TemplateCode = "SMS_203190161"
	} else if sendType == 4 { //修改账号信息
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
	m.SetHeader("From", "coolask@163.com")
	m.SetHeader("To", to...)
	m.SetHeader("Subject", title)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(conf.EmailConf.Host, conf.EmailConf.Port, conf.EmailConf.Username, conf.EmailConf.Password)
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
