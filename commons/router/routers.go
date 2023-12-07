package routers

import (
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
