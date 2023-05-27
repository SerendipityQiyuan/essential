package RoleController

import (
	"awesomeProject1/common"
	"awesomeProject1/model"
	"awesomeProject1/response"
	"awesomeProject1/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	//获取数据
	var requestRole = model.Role{}
	err := ctx.Bind(&requestRole)
	if err != nil {
		return
	}
	password := requestRole.Password
	telephone := requestRole.Telephone
	//password := ctx.PostForm("password")
	//telephone := ctx.PostForm("telephone")

	//数据验证
	//手机号为11位
	if len(telephone) != 11 {
		//StatusUnprocessableEntity -> 状态码：422
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位!")
		return
	}
	//密码不能少于六位
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位!")
		return
	}

	//验证手机号是否存在
	role := model.Role{}
	err = DB.Where(
		"telephone = ?",
		telephone,
	).First(&role).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(role.Password, password)
	//手机号存在 验证密码
	passwordIsTrue := utils.VerifyPassword(
		[]byte(role.Password),
		[]byte(password),
	)
	if passwordIsTrue == false {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	//验证通过 发放token
	token, err := common.ReleaseRoleToken(role)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}

	//返回结果
	response.Response(ctx, http.StatusOK, 200, gin.H{"token": token}, "token已发放")
}
