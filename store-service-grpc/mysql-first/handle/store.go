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

	var databaseList []common.Database
	var database common.Database
	education := service.GetColumnNameList(dao.Education)
	database.DatabaseName = "df_education"
	database.TableList = education
	databaseList = append(databaseList, database)

	library := service.GetColumnNameList(dao.Library)
	database.DatabaseName = "df_library"
	database.TableList = library
	databaseList = append(databaseList, database)

	ctx.JSON(http.StatusOK, databaseList)
}

type Result struct {
	Data []map[string]string
}

// GetColumnData 根据字段名获取数据
func (*StoreHandle) GetColumnData(ctx *gin.Context) {
	var database common.Database

	ctx.BindJSON(&database)
	log.Printf("tableList => %+v", database)
	list := service.GetDataByColumnList(database.DatabaseName, database.TableList)

	ctx.JSON(http.StatusOK, list)
}
