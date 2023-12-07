package router

import (
	routers "commons/router"
	"github.com/gin-gonic/gin"
)
import store "store-service/router/store"

func InitRouter(r *gin.Engine) {
	rg := routers.New()
	rg.Route(&store.StoreRouter{}, r)
}
