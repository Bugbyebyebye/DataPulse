package main

import (
	"auth-service/router"
	"commons/config"
	app "commons/config/app"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//设置路由
	router.InitRouter(r)

	//grpc := router.RegisterGrpc() //注册Grpc
	//router.EtcdServerRegister() //grpc服务注册到etcd
	//stop := func() {
	//	grpc.Stop()
	//}

	//动态启停
	app.Run(r, config.Conf.SC.Name, config.Conf.SC.Addr, nil)
}
