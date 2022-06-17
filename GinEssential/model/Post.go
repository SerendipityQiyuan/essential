package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID         uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	UserId     uint      `json:"user_id" gorm:"not null"`
	CategoryId uint      `json:"category_id" gorm:"not null"`
	Category   *Category
	Title      string `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg    string `json:"head_img" gorm:"type:varchar(255)"`
	Content    string `json:"content" gorm:"type:text"`
	CreatedAt  Time   `json:"created_at" gorm:"autoCreateTime;type:timestamp"`
	UpdatedAt  Time   `json:"updated_at" gorm:"autoUpdateTime;type:timestamp"`
}

func (base *Post) BeforeCreate(scope *gorm.DB) (err error) {
	base.ID = uuid.NewV4()
	return
}
