package router

import (
	"gateway-service/handle"
	"github.com/gin-gonic/gin"
)

type Filter struct {
}

// Route 网关路由匹配
func (*Filter) Route(r *gin.Engine) {
	//log-service 服务
	r.Any("/auth/*path", handle.ProxyHandler)
	r.Any("/log/*path", handle.ProxyHandler)
	r.Any("/store/*path", handle.ProxyHandler)
	r.Any("/task/*path", handle.ProxyHandler)
}
