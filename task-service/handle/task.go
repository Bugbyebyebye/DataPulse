package handle

import (
	"commons/result"
	"github.com/gin-gonic/gin"
)

var res result.Result

func (*TaskHandle) Default(ctx *gin.Context) {

	ctx.JSON(200, res.Success("task服务"))
}
