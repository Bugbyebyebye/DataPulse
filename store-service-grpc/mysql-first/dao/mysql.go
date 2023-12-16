package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	Education *gorm.DB
	Library   *gorm.DB
	err       error
)

func init() {
	//教育数据库
	Education, err = gorm.Open(mysql.Open("root:maojiukeai1412@tcp(222.186.50.126:20010)/df_education?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Printf("mysql error => %s", err)
	}
	if Education.Error != nil {
		log.Printf("Education error => %s", Education.Error)
	}

	//图书馆数据库
	Library, err = gorm.Open(mysql.Open("root:maojiukeai1412@tcp(222.186.50.126:20010)/df_library?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Printf("mysql error => %s", err)
	}
	if Library.Error != nil {
		log.Printf("Education error => %s", Library.Error)
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
