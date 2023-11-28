package auth

import (
	"auth-service/handle"
	"github.com/gin-gonic/gin"
)

type RouterAuth struct {
}

func (*RouterAuth) Route(r *gin.Engine) {
	auth := handle.New()
	//登录注册
	r.POST("/login", auth.UserLogin)
	r.POST("/register", auth.UserRegister)
	//个人信息
	r.GET("/getInfo", auth.GetUserInfo)
	r.POST("/setInfo", auth.SetUserInfo)
}
