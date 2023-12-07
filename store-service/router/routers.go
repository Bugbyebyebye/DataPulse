package router

import (
	routers "commons/router"
	"github.com/gin-gonic/gin"
	"store-service/client"
)
import store "store-service/router/store"

func InitRouter(r *gin.Engine) {
	client.InitDataBaseClient()
	rg := routers.New()
	rg.Route(&store.StoreRouter{}, r)
}
