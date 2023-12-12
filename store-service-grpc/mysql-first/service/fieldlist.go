package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// GetColumnNameList  获取数据表的列名
func GetColumnNameList(db *gorm.DB) []Table {
	var tableList []Table
	var table Table
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
		var info []TableInfo
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
