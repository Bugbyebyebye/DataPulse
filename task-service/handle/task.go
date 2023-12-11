package handle

import (
	"commons/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

var res result.Result

func (*TaskHandle) Default(ctx *gin.Context) {

	ctx.JSON(200, res.Success("task服务"))
}

func (*TaskHandle) CreateApi(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, res.Success("服务生成成功！"))
}
func (*TaskHandle) RunDocker(ctx *gin.Context) {
	port := ctx.Query("port")
	name := ctx.Query("name")
	log.Printf("port %s name %s", port, name)
	go api.RunDocker(port, name)

	ctx.JSON(http.StatusOK, res.Success("服务生成成功！"))
}
