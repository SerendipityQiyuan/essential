package middleware

import (
	"awesomeProject1/common"
	"awesomeProject1/model"
	"awesomeProject1/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// RoleMiddleware 保护用户信息接口
func RoleMiddleware() gin.HandlerFunc {
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

		//验证通过后获取claims中的roleID
		userId := claims.UserId
		DB := common.GetDB()
		var role model.Role
		err = DB.First(&role, userId).Error
		if err != nil {
			panic(err)
		}

		//角色存在，将role信息写入上下文
		ctx.Set("role", role)

		ctx.Next()
	}
}
