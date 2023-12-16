package handle

import (
	"github.com/gin-gonic/gin"
	"log"
	"mysql-second/dao"
	"mysql-second/service"
	"net/http"
)

func (*StoreHandle) GetAllColumnNameList(ctx *gin.Context) {
	var databaseList []service.Database
	var database service.Database
	education := service.GetColumnNameList(dao.Department)
	database.DatabaseName = "df_department"
	database.TableList = education
	databaseList = append(databaseList, database)

	ctx.JSON(http.StatusOK, databaseList)
}

// GetColumnData 查询字段对应的数据
func (*StoreHandle) GetColumnData(ctx *gin.Context) {

	var result []map[string]interface{}
	tableName := ctx.PostForm("table_name")
	err := dao.Department.Table(tableName).Find(&result).Error
	if err != nil {
		log.Printf("err => %s", err)
	}

	ctx.JSON(http.StatusOK, result)
}
