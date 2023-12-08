package router

import (
	"github.com/gin-gonic/gin"
	"task-service/handle"
)

type TaskRouter struct {
}

func (*TaskRouter) Route(r *gin.Engine) {
	task := handle.New()
	//功能路由
	r.GET("/default", task.Default)
	r.GET("/create", task.CreateApi)
}
