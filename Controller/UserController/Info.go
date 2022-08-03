package UserController

import (
	"awesomeProject1/common"
	"awesomeProject1/dto"
	"awesomeProject1/model"
	"awesomeProject1/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func DefaultInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))}, "信息返回成功")
}

func OrderUserInfo(ctx *gin.Context) {
	db := common.GetDB()
	uid := ctx.Query("uid")
	var user model.User
	err := db.Where("id = ?", uid).First(&user).Error
	if err != nil {
		response.Fail(ctx, nil, "用户信息返回失败")
		return
	}
	fmt.Println("uid: ", uid)
	response.Success(ctx, gin.H{"user": dto.ToUserDto(user)}, "信息返回成功")
}
