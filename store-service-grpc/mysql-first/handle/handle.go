package handle

import (
	"mysql-first/dao"
)

// Database databaseList
type Database struct {
	DatabaseName string      `json:"database_name"`
	TableList    interface{} `json:"table_list"`
}

// GetColumnNameList 获取数据库表字段列表
func GetColumnNameList() interface{} {
	var databaseList []Database
	var database Database
	education := dao.GetColumnNameList(dao.Education)
	database.DatabaseName = "df_education"
	database.TableList = education
	databaseList = append(databaseList, database)

	library := dao.GetColumnNameList(dao.Library)
	database.DatabaseName = "df_education"
	database.TableList = library
	databaseList = append(databaseList, database)

	return databaseList
}
