package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          string `json:"name" gorm:"type:varchar(20);not null"`
	Password      string `json:"password" gorm:"type:varchar(255);not null"`
	Telephone     string `json:"telephone" gorm:"type:varchar(11);not null;unique"`
	Sex           string `json:"sex" gorm:"type:char(2);default:'未知'"`
	Age           uint   `json:"age" gorm:"type:int"`
	Introduce     string `json:"introduce" gorm:"type:varchar(255)"`
	PortraitImage string `json:"portrait_image" gorm:"type:varchar(255)"`
}
