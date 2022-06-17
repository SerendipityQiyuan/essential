package model

type Category struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" gorm:"type:varchar(50);not null;unique"`
	CreateAt Time   `json:"create_at" gorm:"AutoCreateTime;type:timestamp"`
	UpdateAt Time   `json:"update_at" gorm:"AutoUpdateTime;type:timestamp"`
}
