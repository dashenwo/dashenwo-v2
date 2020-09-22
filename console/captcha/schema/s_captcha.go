package schema

// 用户信息字段（对外展示）
type Captcha struct {
	// 数据库物理编号
	ID string `json:"id"`
	// 手机号
	Recipient string `json:"recipient"`
	// 验证码
	Code string `json:"code"`
	// 验证码类型
	Type int `json:"type"`
}
