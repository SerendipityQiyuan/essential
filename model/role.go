package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Role struct {
	Id        uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Authority string    `json:"authority" gorm:"type:varchar(255);not null;default:normal"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	Telephone string    `json:"telephone" gorm:"type:varchar(11);not null"`
}

func (base *Role) BeforeCreate(scope *gorm.DB) (err error) {
	base.Id = uuid.NewV4()
	return
}
