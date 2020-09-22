package model

type Captcha struct {
	// 物理id，自增，主键
	ID string `gorm:"column:id;primaryKey;type:varchar(20)"`
	// 手机号
	Recipient string `gorm:"column:recipient;type:varchar(50)"`
	// 验证码
	Code string `gorm:"column:code;type:varchar(8)"`
	// 类型
	Type int `gorm:"column:type;type:int(2)"`
	Model
}
