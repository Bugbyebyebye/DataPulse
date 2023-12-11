package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var (
	Department *gorm.DB
	err        error
)

func init() {
	//教育数据库
	Department, err = gorm.Open("mysql", "root:maojiukeai1412@tcp(222.186.50.126:20085)/df_department?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("mysql error => %s", err)
	}
	if Department.Error != nil {
		log.Printf("Education error => %s", Department.Error)
	}
	Department.DB().SetMaxIdleConns(10)
	Department.DB().SetMaxOpenConns(100)
	Department.DB().SetConnMaxLifetime(time.Hour)
}
