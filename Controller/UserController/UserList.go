package UserController

import (
	"awesomeProject1/common"
	"awesomeProject1/model"
	"awesomeProject1/response"
	"fmt"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	GetUserList(ctx *gin.Context)
}
type UserController struct {
	DB *gorm.DB
}

func NewUserListController() IUserController {
	db := common.GetDB()
	return UserController{DB: db}
}

func (u UserController) GetUserList(ctx *gin.Context) {
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	var users []model.User
	var total int64

	err := u.DB.Order("created_at desc").
		Select("created_at", "updated_at", "id", "name", "telephone", "sex", "age", "introduce", "portrait_image").
		Limit(pageNum - 1).Offset(pageSize * (pageNum - 1)).Find(&users).Error
	//	err := u.DB.Order("created_at desc").
	//		Limit(pageSize).Offset(pageSize * (pageNum - 1)).Find(&users).Error
	if err != nil {
		response.Fail(ctx, gin.H{"data": err}, "用户列表获取失败")
		return
	}
	u.DB.Model(&users).Count(&total)
	fmt.Println(users)
	response.Success(ctx, gin.H{"data": users, "total": total}, "用户列表获取成功")
}
