package schema

// 用户信息字段（对外展示）
type Account struct {
	// 数据库物理编号
	ID string `json:"id"`
	// 昵称
	Nickname string `json:"nickname"`
	// 密码
	Password string `json:"password"`
	// 手机号
	Phone string `json:"phone"`
	// 邮箱
	Email string `json:"email"`
	// 性别
	Sex int `json:"sex"`
	// 生日
	Birthday int64 `json:"birthday"`
	// 签名
	UserSign string `json:"user_sign"`
	// 粉丝数量
	FansNum int64 `json:"fans_num"`
	// 头像地址
	AvatarHd string `json:"avatar_hd"`
	// 关注数量
	FollowNum int64 `json:"follow_num"`
	// 获得的点赞数量
	LoveNum int64 `json:"love_num"`
	// 注册时间
	RegisterTime int64 `json:"register_time"`
}
