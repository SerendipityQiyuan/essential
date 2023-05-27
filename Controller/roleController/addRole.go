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

func RegisterRole(ctx *gin.Context) {
	DB := common.GetDB()
	//获取数据

	//使用map获取请求的参数
	//var requestMap = make(map[string]string)
	//json.NewDecoder(ctx.Request.Body).Decode(&requestMap)

	//使用结构体获取参数
	var requestRole = model.Role{}
	//json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	err := ctx.Bind(&requestRole)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	//name := ctx.PostForm("name")
	//password := ctx.PostForm("password")
	//telephone := ctx.PostForm("telephone")
	name := requestRole.Name
	password := requestRole.Password
	telephone := requestRole.Telephone
	fmt.Println(
		"name:", name,
		"password:", password,
		"telephone:", telephone,
	)

	//数据验证
	//手机号为11位
	if len(telephone) != 11 {
		//StatusUnprocessableEntity -> 状态码：422
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位!")
		return
	}
	//密码不能少于六位
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位!")
		return
	}
	//名称验证 若为空则给一个随机十位的字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
		//fmt.Println("name:", name)
	}

	//验证手机号是否存在
	if utils.IsTelephoneExistRole(DB, telephone) {
		//如果为真 则用户已经存在
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}
	//对用户密码进行哈希加密
	password, err = utils.SecretPassword(password)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
	}

	//创建用户
	newRole := model.Role{
		Name:      name,
		Password:  password,
		Telephone: telephone,
	}
	err = DB.Create(&newRole).Error
	if err != nil {
		panic(err)
	}

	//验证通过 发放token
	token, err := common.ReleaseRoleToken(newRole)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}

	//返回结果
	response.Response(ctx, http.StatusOK, 200, gin.H{"token": token}, "注册成功")
}
