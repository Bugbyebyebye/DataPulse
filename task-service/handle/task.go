package handle

import (
	"commons/result"
	"github.com/gin-gonic/gin"
	"log"
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

func (*TaskHandle) ApiData(ctx *gin.Context) {
	url := ctx.Request.URL
	log.Printf("url => %s", url)
	param := ctx.Param("path")
	log.Printf("param => %s", param)
	ctx.JSON(http.StatusOK, res.Success(url))
}
