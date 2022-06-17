package UserController

import (
	"awesomeProject1/common"
	"awesomeProject1/model"
	"awesomeProject1/response"
	"awesomeProject1/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Login(ctx *gin.Context) {
	DB := common.InitDB()
	//获取数据
	var requestUser = model.User{}
	err := ctx.Bind(&requestUser)
	if err != nil {
		return
	}
	password := requestUser.Password
	telephone := requestUser.Telephone
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
	user := model.User{}
	DB.Where(
		"telephone = ?",
		telephone,
	).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	//手机号存在 验证密码
	isTrue := utils.VerifyPassword(
		[]byte(user.Password),
		[]byte(password),
	)
	if isTrue == false {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	//验证通过 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}

	//返回结果
	response.Response(ctx, http.StatusOK, 200, gin.H{"token": token}, "token已发放")
}
