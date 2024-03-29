package router

import (
	"github.com/gin-gonic/gin"
	"mysql-first/handle"
)

type StoreRouter struct {
}

func (*StoreRouter) Route(r *gin.Engine) {
	store := handle.New()
	//获取数据库信息
	r.GET("/getInfo", store.GetAllColumnNameList)
	//获取字段数据
	r.POST("/getColumnData", store.GetColumnData)
}
