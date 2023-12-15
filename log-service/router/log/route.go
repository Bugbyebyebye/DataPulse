package router

import (
	"github.com/gin-gonic/gin"
	"log-service/handle"
)

type LogRouter struct {
}

func (*LogRouter) Route(r *gin.Engine) {
	log := handle.New()
	//Todo
	r.GET("/getLog", log.GetLogInfo)
	//前端获取路由
	r.POST("/deleteuserlogs", log.DeleteUserLogs)
	r.POST("/deleteonelogs", log.DeleteOneLogs)
	r.POST("/getuserlogs", log.GetUserLogs)
	//断点推送路由
	r.POST("/recorduserlog", log.Logging)
}
