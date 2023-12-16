package dao

import (
	"commons/util"
	"fmt"
	"gorm.io/gorm"
	"log"
	"strings"
)

//创建数据表的相关函数

// CreateTableBySQL 通过表名和字段列表创建表
func CreateTableBySQL(db *gorm.DB, tableName string, columns []string) {
	log.Printf("createTableBySQL => tableName %s columns => %s", tableName, columns)
	var createTableSQL string
	createTableSQL = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s INT AUTO_INCREMENT PRIMARY KEY", tableName, "id")
	for _, column := range columns {
		if column[len(column)-2:] == "id" || column == "state" {
			createTableSQL += fmt.Sprintf(", %s INT ", column)
		} else if column[len(column)-4:] == "time" {
			createTableSQL += fmt.Sprintf(", %s BIGINT(11) ", column)
		} else {
			createTableSQL += fmt.Sprintf(", %s VARCHAR(255) ", column)
		}
	}
	createTableSQL += ");"
	db.Exec(createTableSQL)

	fmt.Println("数据表创建成功")
}

// AlertTableBySQL 向已经存在的表中追加字段
func AlertTableBySQL(db *gorm.DB, tableName string, columns []string) {
	log.Printf("tableName => %s columns => %s", tableName, columns)
	var alterTableSQL string
	var columnSQL []string
	columnExists := GetAllColumnName(tableName, db)
	for _, column := range columns {
		if util.In(column, columnExists) {
			continue
		}
		if column[len(column)-2:] == "id" || column == "state" {
			columnSQL = append(columnSQL, fmt.Sprintf("ADD COLUMN %s INT", column))
		} else if column[len(column)-4:] == "time" {
			columnSQL = append(columnSQL, fmt.Sprintf("ADD COLUMN %s BIGINT(11)", column))
		} else {
			columnSQL = append(columnSQL, fmt.Sprintf("ADD COLUMN %s VARCHAR(255)", column))
		}
	}
	alterTableSQL = strings.Join(columnSQL, ",")

	alterTableSQL = fmt.Sprintf("ALTER TABLE %s %s;", tableName, alterTableSQL)
	log.Printf("alterTableSQL => %s", alterTableSQL)

	db.Exec(alterTableSQL)

	fmt.Println("数据表字段追加成功")
}
