package handle

import (
	"github.com/gin-gonic/gin"
	"log"
	"mongodb-first/common"
	"mongodb-first/dao"
	"mongodb-first/service"
	"net/http"
)

// GetDatabaseColumnNameList 获取数据表信息
func (*StoreHandle) GetDatabaseColumnNameList(ctx *gin.Context) {
	//传入连接名
	article := service.GetColumnNameList(dao.Article)
	log.Printf("article => %+v", article)
	ctx.JSON(http.StatusOK, article)
}

// GetColumnData 获取指定字段数据
func (*StoreHandle) GetColumnData(ctx *gin.Context) {
	var table common.Table

	err := ctx.BindJSON(&table)
	if err != nil {
		log.Printf("err => %s", err)
	}
	result := service.GetColumnData(table)
	log.Printf("result => %+v", result)

	ctx.JSON(http.StatusOK, result)
}
