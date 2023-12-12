package handle

import (
	"github.com/gin-gonic/gin"
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
