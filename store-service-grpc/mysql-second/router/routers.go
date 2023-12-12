package router

import (
	routers "commons/router"
	"github.com/gin-gonic/gin"
	router "mysql-second/router/store"
)

func InitRouter(r *gin.Engine) {
	rg := routers.New()
	rg.Route(&router.StoreRouter{}, r)
}
