package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

//store-service 服务自己的配置
//数据库配置等

var (
	System     *gorm.DB
	Warehouse  *gorm.DB
	Warehouse2 *gorm.DB
	err        error
)

func init() {
	//系统数据库
	System, err = gorm.Open(mysql.Open("root:maojiukeai1412@tcp(222.186.50.126:20134)/df_system?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Printf("mysql error => %s", err)
	}
	if System.Error != nil {
		log.Printf("System error => %s", System.Error)
	}

	//数据仓库 数据库
	Warehouse, err = gorm.Open(mysql.Open("root:maojiukeai1412@tcp(222.186.50.126:20134)/df_warehouse?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Printf("mysql error => %s", err)
	}
	if Warehouse.Error != nil {
		log.Printf("Warehouse error => %s", Warehouse.Error)
	}

	Warehouse2, err = gorm.Open(mysql.Open("root:maojiukeai1412@tcp(222.186.50.126:20134)/df_warehouse2?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Printf("mysql error => %s", err)
	}
	if Warehouse2.Error != nil {
		log.Printf("Warehouse2 error => %s", Warehouse2.Error)
	}
}

// GetDbByDatabaseName 根据数据库名 获取数据库操作指针
func GetDbByDatabaseName(databaseName string) *gorm.DB {
	if databaseName == "df_warehouse" {
		return Warehouse
	} else if databaseName == "df_warehouse2" {
		return Warehouse2
	}
	return nil
}
