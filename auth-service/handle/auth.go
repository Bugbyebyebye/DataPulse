package handle

import (
	"auth-service/config"
	"auth-service/dao"
	"commons/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerUser struct {
	cache config.Cache
}

func New() *HandlerUser {
	return &HandlerUser{
		cache: dao.Rc,
	}
}

// UserLogin 用户登录
func (h *HandlerUser) UserLogin(ctx *gin.Context) {
	res := &result.Result{}

	//1.获取参数

	//2.校验参数

	//3.生成验证码

	//4.调用短信平台 放入go协程
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	log.Println("短信平台调用成功，发送短信")
	//	//5.存储验证码 redis
	//
	//	err := h.cache.Put(ctx, "REGISTER_"+mobile, code, 15*time.Minute)
	//	if err != nil {
	//		log.Printf("将手机号存入redis失败：err:%s", err)
	//	} else {
	//		log.Printf("将手机号存入redis成功：REGISTER_%s:%s", mobile, code)
	//	}
	//}()

	ctx.JSON(http.StatusOK, res.Success("登录成功"))
}

// UserRegister 用户注册
func (h *HandlerUser) UserRegister(ctx *gin.Context) {
	res := &result.Result{}

	ctx.JSON(http.StatusOK, res.Success("注册成功"))
}

// GetUserInfo 获取个人中心信息
func (h *HandlerUser) GetUserInfo(ctx *gin.Context) {

}

// SetUserInfo 设置个人中心信息
func (h *HandlerUser) SetUserInfo(ctx *gin.Context) {

}
