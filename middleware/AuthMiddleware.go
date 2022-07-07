package middleware

import (
	"awesomeProject1/common"
	"awesomeProject1/model"
	"awesomeProject1/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// AuthMiddleware 保护用户信息接口
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		//验证token格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "token格式错误")
			//抛弃此次请求
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "token错误")
			//抛弃此次请求
			ctx.Abort()
			return
		}

		//验证通过后获取claims中的userID
		userId := claims.UserId
		DB := common.InitDB()
		var user model.User
		DB.First(&user, userId)

		//用户不存在
		if user.ID == 0 {
			log.Println("userID:", user.ID)
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "用户不存在")
			//抛弃此次请求
			ctx.Abort()
			return
		}

		//用户存在，将user信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}
