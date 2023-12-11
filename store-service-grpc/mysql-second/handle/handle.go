package handle

import "mysql-second/dao"

// Database databaseList
type Database struct {
	DatabaseName string      `json:"database_name"`
	TableList    interface{} `json:"table_list"`
}

// GetColumnNameList 获取数据库表字段列表
func GetColumnNameList() interface{} {
	var databaseList []Database
	var database Database
	education := dao.GetColumnNameList(dao.Department)
	database.DatabaseName = "df_department"
	database.TableList = education
	databaseList = append(databaseList, database)

	return databaseList
}
