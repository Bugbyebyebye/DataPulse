package router

import (
	authClient "gateway-service/client"
	"gateway-service/config"
	router "gateway-service/router/gateway"
	"github.com/gin-gonic/gin"
)

// Router 路由接口
type Router interface {
	Route(r *gin.Engine)
}

// RegisterRouter 注册路由初始化方法
type RegisterRouter struct {
}

func New() *RegisterRouter {
	return &RegisterRouter{}
}

// Route 路由封装类
func (*RegisterRouter) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

func InitRouter(r *gin.Engine) {
	authClient.InitAuthClient()
	//全局跨域中间件
	r.Use(config.Cors())
	rg := New()
	rg.Route(&router.Filter{}, r)
}
