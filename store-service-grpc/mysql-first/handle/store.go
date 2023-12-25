package handle

import (
	"github.com/gin-gonic/gin"
	"log"
	"mysql-first/common"
	"mysql-first/dao"
	"mysql-first/service"
	"net/http"
)

// GetAllColumnNameList 获取mysql1数据库中的所有字段名
func (*StoreHandle) GetAllColumnNameList(ctx *gin.Context) {
	var result []common.Table
	education := service.GetColumnNameList(dao.Education, "df_education")

	library := service.GetColumnNameList(dao.Library, "df_library")

	result = append(education, library...)
	ctx.JSON(http.StatusOK, result)
}

type Result struct {
	Data []map[string]string
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
