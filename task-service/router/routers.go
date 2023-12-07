package router

import (
	routers "commons/router"
	"github.com/gin-gonic/gin"
	task "task-service/router/task"
)

func InitRouter(r *gin.Engine) {
	rg := routers.New()
	rg.Route(&task.TaskRouter{}, r)
}
