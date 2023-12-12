package router

import (
	"github.com/gin-gonic/gin"
	"mysql-second/handle"
)

type StoreRouter struct {
}

func (*StoreRouter) Route(r *gin.Engine) {
	store := handle.New()
	//获取数据库信息
	r.GET("/getInfo", store.GetAllColumnNameList)
}
