package main

import (
	"commons/config"
	"commons/config/app"
	"github.com/gin-gonic/gin"
	"mysql-first/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	app.Run(r, config.Conf.SC.Name, config.Conf.SC.Addr, nil)
}
