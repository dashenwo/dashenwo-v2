package model

type Account struct {
	// 物理id，自增，主键
	ID string `gorm:"column:id;primaryKey;type:varchar(20)"`
	// 用户昵称,唯一键
	Nickname string `gorm:"column:nickname;unique;type:varchar(50)"`
	// 密码
	Password string `json:"-" gorm:"column:password;type:varchar(100)"`
	// 密码生成盐值
	Salt string `json:"-" gorm:"column:salt;type:varchar(8)"`
	// 手机号
	Phone string `gorm:"column:phone;type:varchar(11)"`
	// 邮箱
	Email string `gorm:"column:email;type:varchar(50)"`
	// 性别
	Sex int `gorm:"column:sex;type:int(2)"`
	// 生日
	Birthday int64 `gorm:"column:birthday;type:bigint(13)"`
	// 签名
	UserSign string `gorm:"column:user_sign;type:varchar(100)"`
	// 头像地址
	AvatarHd string `gorm:"column:avatarHd;type:varchar(100)"`
	// 粉丝数量
	FansNum int64 `gorm:"column:fans_num;type:bigint(13)"`
	// 关注数量
	FollowNum int64 `gorm:"column:follow_num;type:bigint(13)"`
	// 获得点赞数量
	LoveNum int64 `gorm:"column:love_num;type:bigint(13)"`
	// 注册时间
	RegisterTime int64 `gorm:"column:register_time;type:bigint(13)"`
	Model
}
