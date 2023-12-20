package router

import (
	"github.com/gin-gonic/gin"
	"task-service/handle"
)

type TaskRouter struct {
}

func (*TaskRouter) Route(r *gin.Engine) {
	task := handle.New()

	r.GET("/docker/run", task.RunDocker)
	r.GET("/docker/stop", task.StopDocker)
	r.GET("/docker/restart", task.RestartDocker)
}
