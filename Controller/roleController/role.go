package RoleController

import (
	"awesomeProject1/common"
	"awesomeProject1/model"
	"awesomeProject1/response"
	"awesomeProject1/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IRoleController interface {
	CreateRole(ctx *gin.Context)
	ChangeRole(ctx *gin.Context)
	DeleteRole(ctx *gin.Context)
}

type UserController struct {
	DB *gorm.DB
}

func NewRoleController() IRoleController {
	db := common.GetDB()
	err := db.AutoMigrate(model.Role{})
	if err != nil {
		panic(err)
	}
	return UserController{DB: db}
}

func (u UserController) CreateRole(ctx *gin.Context) {
	//TODO implement me
	//获取当前用户id

	//userId := ctx.("uid")
	//var roleModel model.Role
	//
	//type getAuthority struct{ Authority string }

	//var userAuthority getAuthority
	//u.DB.Model(&roleModel).Select("authority").Where("id = ?", userId).First(&userAuthority)

	//判断当前用户id权限
	userAuthority, _ := ctx.Get("role")
	if userAuthority.(model.Role).Authority != "administrator" {
		response.Fail(ctx, nil, "权限不足")
		return
	}

	fmt.Println(userAuthority.(model.Role).Authority)

	type requestCreateRole struct {
		Name      string
		Authority string
		Telephone string
	}

	var requestRole requestCreateRole

	err := ctx.ShouldBind(&requestRole)
	role := model.Role{
		Name:      requestRole.Name,
		Authority: requestRole.Authority,
		Telephone: requestRole.Telephone,
	}

	if role.Password == "" {
		role.Password, err = utils.SecretPassword("123456")
		if err != nil {
			panic(err)
		}
	}

	if err != nil {
		panic(err)
	}
	err = u.DB.Create(&role).Error
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"data": role}, "角色创建成功")
}

func (u UserController) ChangeRole(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u UserController) DeleteRole(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
