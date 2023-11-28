package dao

import (
	srv "commons/config"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open("mysql", srv.Conf.MC.Name+":"+srv.Conf.MC.Password+"@tcp("+srv.Conf.MC.Host+")"+"/go_db"+"?charset=utf8")
	if err != nil {
		log.Printf("mysql error: %s", err.Error())
	}

	if Db.Error != nil {
		log.Printf("datebase error: %s", Db.Error)
	}

	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetConnMaxLifetime(time.Hour)
}
