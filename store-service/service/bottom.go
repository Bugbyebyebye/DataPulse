package service

import (
	"github.com/carlmjohnson/requests"
	"github.com/gin-gonic/gin"
	"log"
	"store-service/common"
	"store-service/config"
)

// GetBottomTableInfo 获取底层数据库数据表信息
func GetBottomTableInfo(ctx *gin.Context) []common.Table {
	var result []common.Table

	var first []common.Table
	err := requests.URL(config.Conf.SEVERURL.MYSQLF).
		Path("/getInfo").
		ToJSON(&first).Fetch(ctx)
	if err != nil {
		log.Printf("err => %+v", err)
	}
	var second []common.Table
	err = requests.URL(config.Conf.SEVERURL.MYSQLS).
		Path("/getInfo").
		ToJSON(&second).Fetch(ctx)
	if err != nil {
		log.Printf("err => %+v", err)
	}
	var three []common.Table
	err = requests.URL(config.Conf.SEVERURL.MongoDB).
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
