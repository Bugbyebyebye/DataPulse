package router

import (
	"github.com/gin-gonic/gin"
	"store-service/handle"
)

type StoreRouter struct {
}

func (*StoreRouter) Route(r *gin.Engine) {
	store := handle.New()

	//获取首页统计数据
	r.GET("/getNum", store.GetIndexTableData)
	//获取数据库信息
	//获取底层数据库信息
	r.GET("/getInfo", store.GetBottomDatabaseNameList)
	r.GET("/getTargetInfo", store.GetWarehouseDatabaseNameList)
	//获取数据仓库信息
	r.GET("/getUserTable", store.GetUserDatabaseNameList)
	r.POST("/getTableData", store.GetUserTableData)
	r.POST("/deleteUserTable", store.DeleteWarehouseTable)
	//创建数据仓库新数据表
	r.POST("/createTable", store.CreateTheTable)
	r.POST("/alertTable", store.AlertTable)
}
