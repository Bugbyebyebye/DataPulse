package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	Department *gorm.DB
	err        error
)

func init() {
	//教育数据库
	Department, err = gorm.Open(mysql.Open("root:maojiukeai1412@tcp(222.186.50.126:20085)/df_department?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Printf("mysql error => %s", err)
	}
	if Department.Error != nil {
		log.Printf("Education error => %s", Department.Error)
	}
}

// QueryColumnData 传入字段值获取字段数据
func QueryColumnData(db *gorm.DB, tableName string, columnList []string) []map[string]interface{} {
	var data []map[string]interface{}
	err := db.Table(tableName).Select(columnList).Find(&data).Error
	if err != nil {
		log.Printf("mysql error => %s", err)
	}
	return data
}
