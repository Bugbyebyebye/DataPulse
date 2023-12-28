package dao

import (
	srv "commons/config"
	"database/sql"
	"fmt"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", srv.Conf.MC.Name+":"+srv.Conf.MC.Password+"@tcp("+srv.Conf.MC.Host+")"+"/df_system"+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("数据库连接失败 : %s\n", err.Error())
		//log.Fatal("无法连接到数据库")

	}
	// 确保连接正常
	err = Db.Ping()
	if err != nil {
		fmt.Printf("数据库不健康 : %s\n", err.Error())
		//log.Fatal("数据库不健康")
		//return
	}
	fmt.Printf("数据库已连接")
}
