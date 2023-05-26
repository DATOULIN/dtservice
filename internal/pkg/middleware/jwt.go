package middleware

import (
	"github.com/DATOULIN/dtservice/internal/pkg/app"
	"github.com/DATOULIN/dtservice/internal/pkg/auth"
	"github.com/DATOULIN/dtservice/internal/pkg/errno"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//首先判断当前的请求中获取到jwt_token
		tokenStr := ctx.Request.Header.Get("Authorization")
		response := app.NewResponse(ctx)
		if tokenStr == "" {
			response.ToErrorResponse(errno.UnauthorizedNoTokenError)
			//终止退出
			ctx.Abort()
			return
		}
		clams, err := auth.ParseToken(tokenStr)
		// 判断token是否加入黑名单（注销的时候会把当前的token加入黑名单）
		if err != nil || auth.IsInBlacklist(tokenStr) {
			response.ToErrorResponse(errno.UnauthorizedTokenTimeout)
			ctx.Abort()
			return
		}
		if err != nil {
			response.ToErrorResponse(errno.UnauthorizedTokenError)
			//终止退出
			ctx.Abort()
			return
		}
		ctx.Set("clams", clams)
	}
}
