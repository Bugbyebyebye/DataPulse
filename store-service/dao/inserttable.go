package dao

import (
	"gorm.io/gorm"
	"log"
	"reflect"
)

//向表中初始数据的函数

// InitTableData 传入表名和数据map数组插入数据
func InitTableData(db *gorm.DB, tableName string, bottomData []map[string]interface{}) {
	for _, item := range bottomData {
		err := db.Table(tableName).Create(item).Error
		if err != nil {
			log.Printf("err => %s", err)
		}
	}
}

// UpdateTableData 传入表名和数据map数组更新数据
func UpdateTableData(db *gorm.DB, tableName string, bottomData []map[string]interface{}) {
	for _, item := range bottomData {
		var id string
		keys := reflect.ValueOf(item).MapKeys()
		for _, key := range keys {
			if key.String()[len(key.String())-2:] == "id" {
				id = key.String()
			}
		}
		err := db.Table(tableName).Where(""+id+"=?", item[id]).Updates(item).Error
		if err != nil {
			log.Printf("err => %s", err)
		}
	}
}
