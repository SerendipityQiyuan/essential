package UserController

import (
	"awesomeProject1/dto"
	"awesomeProject1/model"
	"awesomeProject1/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	fmt.Println("userInfo:", user)
	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))}, "信息返回成功")
}
