package router

import (
	"github.com/gin-gonic/gin"
	"task-service/handle"
)

type TaskRouter struct {
}

func (*TaskRouter) Route(r *gin.Engine) {
	task := handle.New()
	//处理服务启停止相关
	r.POST("/docker/run", task.RunDocker)
	r.GET("/docker/stop", task.StopDocker)
	r.GET("/docker/restart", task.RestartDocker)
	//处理用户操作相关
	r.GET("/getuserapilist", task.SearchAPIList)
	r.POST("/deleteapilist", task.DeleteApi)
	r.GET("/searchstate", task.SearchStatusLables)
}
