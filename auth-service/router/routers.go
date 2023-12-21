package router

import (
	auth "auth-service/router/auth"
	routers "commons/router"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	rg := routers.New()
	rg.Route(&auth.AuthRouter{}, r)
}
