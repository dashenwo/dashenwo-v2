package validate

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/micro/go-micro/v2/errors"
	"reflect"
	"strings"
)

// 根据传入的结构体和validate错误信息获取自定义tag的错误消息
func GetMessageByTag(entity interface{}, err validator.FieldError) string {
	//获取资源
	r := reflect.TypeOf(entity).Elem()
	// 获取错误字段
	errField := err.Field()
	// 根据错误字段获取结构体信息
	structField, _ := r.FieldByName(errField)
	// message标签所包含的字符串
	messageStr := structField.Tag.Get("message")
	// 创建一个map
	mesMap := map[string]string{}
	// 分割成数组
	msgArray := strings.Split(messageStr, ",")
	for _, msg := range msgArray {
		// 再次分割成key和value
		kvArray := strings.Split(msg, ":")
		if len(kvArray) > 1 {
			mesMap[kvArray[0]] = kvArray[1]
		} else {
			mesMap[kvArray[0]] = ""
		}
	}

	res := mesMap[err.Tag()]
	// 打印一下信息
	//log.Log(structField)
	//log.Log(structField.Tag.Get("message"))
	//log.Log(structErr)
	//log.Log(mesMap)
	//log.Log(err.Tag())

	return res
}

// 验证数据方法
func Validate(entity interface{}, appId string) error {
	// 验证数据
	zh_ch := zh.New()
	uni := ut.New(zh_ch)
	trans, _ := uni.GetTranslator("zh")
	//验证器
	validate := validator.New()
	//验证器注册翻译器
	_ = zhTranslations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(entity)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			msg := GetMessageByTag(entity, err)
			if msg == "" {
				msg = err.Translate(trans)
			}
			return errors.New(appId, msg, 401)
		}
	}
	return nil
}
