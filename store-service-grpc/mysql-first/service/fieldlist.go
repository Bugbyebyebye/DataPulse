package service

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"mysql-first/common"
	"mysql-first/dao"
)

// GetColumnNameList  获取数据表的列名
func GetColumnNameList(db *gorm.DB) []common.Table {
	var tableList []common.Table
	var table common.Table
	var tables []string
	//获取数据库下的全部数据表名
	rows, err := db.Raw("SHOW TABLES").Rows()
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			fmt.Printf("err => %s\n", err)
		}
		tables = append(tables, name)
	}
	//fmt.Printf("Tables: %s\n", tables)

	//获取数据表中全部字段名
	for _, tableName := range tables {
		var columns []string
		var info []common.TableInfo
		query := "desc " + tableName
		db.Raw(query).Scan(&info)
		for _, v := range info {
			columns = append(columns, v.Field)
		}
		table.TableName = tableName
		table.ColumnList = columns
		tableList = append(tableList, table)
		//fmt.Printf("Table: %s\n Columns: %+v\n", tableName, columns)
	}
	return tableList
}

// GetDataByColumnList 获取数据表中指定字段的数据
func GetDataByColumnList(databaseName string, tableList []common.Table) []map[string]interface{} {
	//根据传入的数据库选择操作数据库的对象
	var db *gorm.DB
	if databaseName == "df_education" {
		db = dao.Education
	} else if databaseName == "df_library" {
		db = dao.Library
	}
	//数据表列表中的table
	table := tableList[0] //列表1
	log.Printf("table => %+v", table)
	log.Printf("column => %+v", table.ColumnList)

	result := dao.QueryColumnData(db, table.TableName, table.ColumnList)
	log.Printf("result => %+v", result)

	return result
}
