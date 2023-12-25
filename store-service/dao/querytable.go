package dao

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"store-service/common"
	"store-service/config"
)

//数据仓库查询操作

// QueryTableData 根据数据库名查询全部数据
func QueryTableData(db *gorm.DB, tableName string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := db.Table(tableName).Find(&result).Error
	return result, err
}

// GetAllTableName 获取数据库中已有表的名称
func GetAllTableName(db *gorm.DB) []string {
	var tables []string
	rows, err := db.Raw("SHOW TABLES").Rows()
	if err != nil {
		log.Printf("err => %s", err.Error())
	}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			fmt.Printf("err => %s\n", err)
		}
		tables = append(tables, name)
	}
	return tables
}

// GetAllColumnName 获取表中已有字段名
func GetAllColumnName(tableName string, db *gorm.DB) []string {
	var columns []string
	var info []common.TableInfo
	query := "desc " + tableName
	db.Raw(query).Scan(&info)
	for _, v := range info {
		columns = append(columns, v.Field)
	}
	return columns
}

// GetDataByColumnList 获取数据表中指定字段的数据
func GetDataByColumnList(table common.Table) []map[string]interface{} {
	databaseName := table.DatabaseName

	//根据传入的数据库选择操作数据库的对象
	var db *gorm.DB
	if databaseName == "df_warehouse" {
		db = config.Warehouse
	} else if databaseName == "df_warehouse2" {
		db = config.Warehouse2
	}
	//数据表列表中的table
	log.Printf("table => %+v", table)
	log.Printf("column => %+v", table.ColumnList)

	result := config.QueryColumnData(db, table.TableName, table.ColumnList)
	log.Printf("result => %+v", result)

	return result
}
