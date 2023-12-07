package main

import (
	"commons/config"
	"commons/config/app"
	"github.com/gin-gonic/gin"
	"task-service/router"
)

func main() {
	r := gin.Default()
	//设置路由
	router.InitRouter(r)
	//动态启停
	app.Run(r, config.Conf.SC.Name, config.Conf.SC.Addr, nil)
}
