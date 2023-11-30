package router

import (
	"github.com/gin-gonic/gin"
	"log-service/handle"
)

type LogRouter struct {
}

func (*LogRouter) Route(r *gin.Engine) {
	log := handle.New()

	r.GET("/getLog", log.GetLogInfo)
}
