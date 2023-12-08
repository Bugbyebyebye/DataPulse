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

	//个人中心
	r.GET("/getInfo", auth.GetUserInfo)
	r.POST("/setInfo", auth.SetUserInfo)
	r.POST("/setAccount", auth.SetAccount)

	//用户管理
	//Certification 用于从token获取信息
	r.GET("/getUserList", Certification(), auth.GetUserList)
	r.POST("/setUserRole", Certification(), auth.SetUserRole)
	r.POST("/setUserAuth", Certification(), auth.SetUserAuthority)
	r.POST("/deleteUser", Certification(), auth.DeleteUser)
}
