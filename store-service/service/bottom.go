package service

import (
	"github.com/carlmjohnson/requests"
	"github.com/gin-gonic/gin"
	"log"
	"store-service/common"
)

// GetBottomTableInfo 获取底层数据库数据表信息
func GetBottomTableInfo(ctx *gin.Context) []common.Table {
	var result []common.Table

	var first []common.Table
	err := requests.URL("http://mysql-first:8085").
		Path("/getInfo").
		ToJSON(&first).Fetch(ctx)
	if err != nil {
		log.Printf("err => %+v", err)
	}
	var second []common.Table
	err = requests.URL("http://mysql-second:8086").
		Path("/getInfo").
		ToJSON(&second).Fetch(ctx)
	if err != nil {
		log.Printf("err => %+v", err)
	}
	var three []common.Table
	err = requests.URL("http://mongodb-first:8087").
		Path("/getInfo").
		ToJSON(&three).Fetch(ctx)
	if err != nil {
		log.Printf("err => %+v", err)
	}

	result = append(result, first...)
	result = append(result, second...)
	result = append(result, three...)

	return result
}
