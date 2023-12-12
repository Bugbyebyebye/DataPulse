package service

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"mysql-first/dao"
	"strings"
)

func GetAllColumnData(param []interface{}) interface{} {
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

// GetColumnData 获取数据表中指定字段的数据
func GetColumnData(databaseName string, tableName string, columns []string) {
	var db *gorm.DB
	if databaseName == "df_education" {
		db = dao.Education
	} else if databaseName == "df_library" {
		db = dao.Library
	}

	var result []interface{}
	err := db.Table(tableName).Select(columns).Find(&result).Error
	if err != nil {
		log.Printf("err => %s", err)
	}

	for _, item := range result {
		log.Printf("item => %+v\n", item)
	}
}
