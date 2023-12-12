package handle

import (
	"github.com/gin-gonic/gin"
	"mysql-first/dao"
	"mysql-second/service"
	"net/http"
)

// GetAllColumnNameList 获取mysql1数据库中的所有字段名
func (*StoreHandle) GetAllColumnNameList(ctx *gin.Context) {

	var databaseList []service.Database
	var database service.Database
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
