package middleware

import (
	"awesomeProject1/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				response.Fail(ctx, nil, fmt.Sprint(err))
				return
			}
		}()
		ctx.Next()
	}
}
