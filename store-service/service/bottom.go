package service

import (
	"github.com/carlmjohnson/requests"
	"github.com/gin-gonic/gin"
	"log"
	"store-service/common"
	"store-service/config"
	"store-service/dao"
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

// GetSourceData 从底层获取数据
func GetSourceData(ctx *gin.Context, table common.Table) []map[string]interface{} {
	var data []map[string]interface{}

	if table.SourceName == "mysql" {
		data = dao.GetDataByColumnList(table)
		//log.Printf("bottom => %+v\n", bottom)
	} else if table.SourceName == "mysql1" {
		err := requests.URL(config.Conf.SEVERURL.MYSQLF).
			Path("/getColumnData").
			BodyJSON(&table).
			ToJSON(&data).
			Fetch(ctx)
		//log.Printf("bottom => %+v\n", bottom)
		if err != nil {
			log.Printf("err => %s", err)
		}
	} else if table.SourceName == "mysql2" {
		err := requests.URL(config.Conf.SEVERURL.MYSQLS).
			Path("/getColumnData").
			BodyJSON(&table).
			ToJSON(&data).
			Fetch(ctx)
		//log.Printf("bottom => %+v", bottom)
		if err != nil {
			log.Printf("err => %s", err)
		}
	} else if table.SourceName == "mongodb1" {
		err := requests.URL(config.Conf.SEVERURL.MongoDB).
			Path("/getColumnData").
			BodyJSON(&table).
			ToJSON(&data).
			Fetch(ctx)
		//log.Printf("bottom => %+v", bottom)
		if err != nil {
			log.Printf("err => %s", err)
		}
	}

	//TODO 数据清洗过滤后输出数据
	//result := CleanData(data)
	//return result
	return data
}

// CleanData 数据清洗
//func CleanData(data []map[string]interface{}) []map[string]interface{} {
//
//}
