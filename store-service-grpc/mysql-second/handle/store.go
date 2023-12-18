package handle

import (
	"github.com/gin-gonic/gin"
	"log"
	"mysql-second/common"
	"mysql-second/dao"
	"mysql-second/service"
	"net/http"
)

func (*StoreHandle) GetAllColumnNameList(ctx *gin.Context) {
	department := service.GetColumnNameList(dao.Department)
	ctx.JSON(http.StatusOK, department)
}

// GetColumnData 根据字段名获取数据
func (*StoreHandle) GetColumnData(ctx *gin.Context) {
	var table common.Table

	err := ctx.BindJSON(&table)
	if err != nil {
		log.Printf("err => %s", err)
	}
	log.Printf("tableList => %+v", table)
	list := service.GetDataByColumnList(table)

	ctx.JSON(http.StatusOK, list)
}
