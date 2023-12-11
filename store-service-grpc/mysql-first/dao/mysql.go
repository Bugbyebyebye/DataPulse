package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var (
	Education *gorm.DB
	Library   *gorm.DB
	err       error
)

func init() {
	//教育数据库
	Education, err = gorm.Open("mysql", "root:maojiukeai1412@tcp(222.186.50.126:20010)/df_education?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("mysql error => %s", err)
	}
	if Education.Error != nil {
		log.Printf("Education error => %s", Education.Error)
	}
	Education.DB().SetMaxIdleConns(10)
	Education.DB().SetMaxOpenConns(100)
	Education.DB().SetConnMaxLifetime(time.Hour)

	//图书馆数据库
	Library, err = gorm.Open("mysql", "root:maojiukeai1412@tcp(222.186.50.126:20010)/df_library?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("mysql error => %s", err)
	}
	if Library.Error != nil {
		log.Printf("Education error => %s", Library.Error)
	}
	Library.DB().SetMaxIdleConns(10)
	Library.DB().SetMaxOpenConns(100)
	Library.DB().SetConnMaxLifetime(time.Hour)
}
