package router

import (
	routers "commons/router"
	authClient "gateway-service/client"
	"gateway-service/config"
	router "gateway-service/router/gateway"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	authClient.InitAuthClient()
	//全局跨域中间件
	r.Use(config.Cors())
	rg := routers.New()
	rg.Route(&router.Filter{}, r)
}
