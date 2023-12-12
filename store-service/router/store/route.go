package router

import (
	"github.com/gin-gonic/gin"
	"store-service/handle"
)

type StoreRouter struct {
}

func (*StoreRouter) Route(r *gin.Engine) {
	store := handle.New()
	//获取数据库信息
	r.GET("/getInfo", store.GetDatabaseColumnNameList)
	//创建数据仓库新数据表
	r.POST("/createTable", store.CreateTable)
	r.POST("/alertTable", store.AlertTable)
}
