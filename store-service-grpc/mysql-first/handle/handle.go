package handle

import (
	"database/sql"
	"fmt"
	"log"
	"mysql-first/dao"
	"strings"
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
	database.DatabaseName = "df_library"
	database.TableList = library
	databaseList = append(databaseList, database)

	return databaseList
}

func GetColumnData(param []interface{}) interface{} {
	results := make([]map[string]string, 0)

	for _, database := range param {
		dbMap, ok := database.(map[string]interface{})
		if !ok {
			continue
		}
		dbName := dbMap["database_name"]
		tableList := dbMap["table_list"].([]interface{})
		for _, table := range tableList {
			tbMap, ok := table.(map[string]interface{})
			if !ok {
				continue
			}
			tbName := tbMap["table_name"].(string)
			columnList := make([]string, 0)
			for _, column := range tbMap["column_list"].([]interface{}) {
				columnList = append(columnList, column.(string))
			}
			log.Printf("databaseName => %s tableName => %s columnlist => %v", dbName, tbName, columnList)

			// 构建SQL查询语句
			query := fmt.Sprintf("SELECT %s FROM `%s`", strings.Join(columnList, ","), tbName)

			// 执行查询并获取结果
			rows, err := dao.Library.DB().Query(query)
			if err != nil {
				panic(err)
			}
			defer rows.Close()

			// 遍历结果

			values := make([]sql.RawBytes, len(columnList))
			scanArgs := make([]interface{}, len(values))
			for i := range values {
				scanArgs[i] = &values[i]
			}
			for rows.Next() {
				err := rows.Scan(scanArgs...)
				if err != nil {
					panic(err)
				}
				result := make(map[string]string)
				for i, col := range columnList {
					result[col] = string(values[i])
				}
				results = append(results, result)
				fmt.Printf("result: %+v\n", result)
			}
			// 检查是否有其他错误
			err = rows.Err()
			if err != nil {
				panic(err)
			}
		}
	}
	return results
}
