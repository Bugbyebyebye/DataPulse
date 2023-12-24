package router

import (
	"mongodb-first/handle"

	"github.com/gin-gonic/gin"
)

type StoreRouter struct {
}

func (*StoreRouter) Route(r *gin.Engine) {
	store := handle.New()
	//获取数据库信息
	r.GET("/getInfo", store.GetDatabaseColumnNameList)
	r.POST("/getColumnData", store.GetColumnData)
	r.GET("/getSchoolWallCleanedInfo", store.GetSchoolWallCleanedInfo)
	r.GET("/getWeChatCleanedInfo", store.GetWeChatCleanedInfo)
}
