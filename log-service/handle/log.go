package handle

import (
	"commons/result"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type LogHandler struct {
}

func New() *LogHandler {
	return &LogHandler{}
}

func (*LogHandler) GetLogInfo(ctx *gin.Context) {
	r := &result.Result{}
	id := ctx.GetHeader("id")
	log.Printf("id => %+v", id)
	ctx.JSON(http.StatusOK, r.Success("查看日志成功！"))
}
