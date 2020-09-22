package model

type Model struct {
	CreatedTime int32 `gorm:"column:created_time;size:36;"`
	UpdatedTime int32 `gorm:"column:updated_time;index"`
	DeletedTime int32 `json:"column:deleted_time;index"`
}
