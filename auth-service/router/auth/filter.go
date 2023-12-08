package router

import (
	"auth-service/util"
	"commons/result"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var res result.Result

// Certification 权限认证中间件
func Certification() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			ctx.JSON(http.StatusOK, res.Fail(4001, "token 为空"))
			ctx.Abort()
		}

		claims, err := util.ParseToken(token)
		if err != nil {
			log.Printf("token 校验错误 err => %s", err)
			ctx.JSON(http.StatusOK, res.Fail(4001, "token 校验错误"))
			ctx.Abort()
		}

		id := claims.Id
		username := claims.Username
		role := claims.Role
		authority := claims.Authority

		if id == 0 && role != "admin" && authority == 0 {
			ctx.JSON(http.StatusOK, res.Fail(4001, "抱歉，权限不够！"))
			return
		}

		ctx.Set("id", id)
		ctx.Set("username", username)
		ctx.Set("role", role)
		ctx.Set("authority", authority)

		ctx.Next()
	}
}
