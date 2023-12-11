package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var (
	System    *gorm.DB
	Warehouse *gorm.DB
	err       error
)

func init() {
	//系统数据库
	System, err = gorm.Open("mysql", "root:maojiukeai1412@tcp(222.186.50.126:20134)/df_system?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("mysql error => %s", err)
	}
	if System.Error != nil {
		log.Printf("Education error => %s", System.Error)
	}
	System.DB().SetMaxIdleConns(10)
	System.DB().SetMaxOpenConns(100)
	System.DB().SetConnMaxLifetime(time.Hour)

	//数据仓库 数据库
	Warehouse, err = gorm.Open("mysql", "root:maojiukeai1412@tcp(222.186.50.126:20134)/df_warehouse?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("mysql error => %s", err)
	}
	if Warehouse.Error != nil {
		log.Printf("Education error => %s", Warehouse.Error)
	}
	Warehouse.DB().SetMaxIdleConns(10)
	Warehouse.DB().SetMaxOpenConns(100)
	Warehouse.DB().SetConnMaxLifetime(time.Hour)
}

// GetAllTableName 获取数据库中已有表的名称
func GetAllTableName(db *gorm.DB) []string {
	var tables []string
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
	return tables
}
