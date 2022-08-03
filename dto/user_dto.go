package dto

import "awesomeProject1/model"

type UserDto struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	Telephone     string `json:"telephone"`
	Sex           string `json:"sex" gorm:"type:char(2)"`
	Age           uint   `json:"age" gorm:"type:int(4)"`
	Introduce     string `json:"introduce" gorm:"type:varchar(555)"`
	PortraitImage string `json:"portrait_image" gorm:"type:varchar(255)"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Id:            user.ID,
		Name:          user.Name,
		Telephone:     user.Telephone,
		Sex:           user.Sex,
		Age:           user.Age,
		Introduce:     user.Introduce,
		PortraitImage: user.PortraitImage,
	}
}
