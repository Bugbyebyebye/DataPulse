package router

import (
	"auth-service/handle"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func (*AuthRouter) Route(r *gin.Engine) {
	auth := handle.New()
	//功能路由
	r.GET("/code", auth.PostEmailCode)
	r.POST("/file", auth.UploadPic)
	//登录注册
	r.POST("/login", auth.UserLogin)
	r.POST("/register", auth.UserRegister)
	//个人信息
	r.GET("/getInfo", auth.GetUserInfo)
	r.POST("/setInfo", auth.SetUserInfo)
	//账号信息相关
	r.POST("/setAccount", auth.SetAccount)

}
