package handle

import (
	"commons/util"
	"fmt"
	"github.com/carlmjohnson/requests"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"store-service/common"
	"store-service/config"
	"store-service/dao"
	"store-service/model"
)

//创建数据表 追加字段

// CreateTheTable 创建数据仓库数据表
func (*StoreHandle) CreateTheTable(ctx *gin.Context) {
	//从token获取用户id
	id := ctx.GetInt("id")

	var jsonParam common.DataSource
	ctx.BindJSON(&jsonParam)
	databaseName := jsonParam.TargetDatabase                     //目标数据库名
	tableName := jsonParam.TargetTable                           //目标表名
	tableType := jsonParam.TargetType                            //目标表类型
	sourceName := jsonParam.FromName                             //数据源名
	databaseList := jsonParam.DatabaseList[0]                    //数据表列表 1
	columns := jsonParam.DatabaseList[0].TableList[0].ColumnList //数据表1字段列表

	//TODO 应该针对多个数据表的字段进行创表和插入数据
	db := config.GetDbByDatabaseName(databaseName)
	log.Printf("db => %+v", db)

	//检查是否有重名表
	tableNameExist := dao.GetAllTableName(db)
	if util.In(tableName, tableNameExist) {
		ctx.JSON(http.StatusOK, res.Fail(4001, "表名已存在，请重新选择"))
		return
	}

	//创建新数据库
	dao.CreateTableBySQL(db, tableName, columns)
	//创建用户数据库关联
	tableId, err := model.InitTable(tableName, databaseName, tableType, len(columns))
	if err != nil {
		log.Printf("err => %s", err)
	}
	//初始化数据表和用户关联
	err = model.InitTableUser(tableId, id)
	if err != nil {
		log.Printf("err => %s", err)
	}
	log.Printf("tableList => %+v", databaseList)

	//根据数据源选择底层数据来源
	var bottomData []map[string]interface{}
	if sourceName == "mysql1" {
		err := requests.URL("http://localhost:8085").
			Path("/getColumnData").
			BodyJSON(&databaseList).
			ToJSON(&bottomData).
			Fetch(ctx)
		//log.Printf("bottomData => %+v", bottomData)
		if err != nil {
			log.Printf("err => %s", err)
		}
	} else if sourceName == "mysql2" {
		err := requests.URL("http://localhost:8086").
			Path("/getColumnData").
			BodyJSON(&databaseList).
			ToJSON(&bottomData).
			Fetch(ctx)
		//log.Printf("bottomData => %+v", bottomData)
		if err != nil {
			log.Printf("err => %s", err)
		}
	}
	//根据底层数据 新增数据
	dao.InitTableData(db, tableName, bottomData)

	fmt.Println("Data inserted successfully.")
	ctx.JSON(http.StatusOK, res.Success(bottomData))
}

// AlertTable 向数据仓库数据表中追加字段
func (*StoreHandle) AlertTable(ctx *gin.Context) {
	var jsonParam common.DataSource
	ctx.BindJSON(&jsonParam)

	databaseName := jsonParam.TargetDatabase
	tableName := jsonParam.TargetTable                           //表明
	sourceName := jsonParam.FromName                             //数据源名
	databaseList := jsonParam.DatabaseList[0]                    //数据表列表 1
	columns := jsonParam.DatabaseList[0].TableList[0].ColumnList //数据表1字段列表
	//TODO 应该针对多个数据表的字段进行创表和插入数据

	//匹配目标数据库
	db := config.GetDbByDatabaseName(databaseName)
	//向数据表中追加新增字段
	dao.AlertTableBySQL(db, tableName, columns)

	var bottomData []map[string]interface{}
	//根据数据源选择底层数据来源
	if sourceName == "mysql1" {
		err := requests.URL("http://localhost:8085").
			Path("/getColumnData").
			BodyJSON(&databaseList).
			ToJSON(&bottomData).
			Fetch(ctx)
		log.Printf("bottomData => %+v", bottomData)
		if err != nil {
			log.Printf("err => %s", err)
		}
	} else if sourceName == "mysql2" {
		err := requests.URL("http://localhost:8086").
			Path("/getColumnData").
			BodyJSON(&databaseList).
			ToJSON(&bottomData).
			Fetch(ctx)
		log.Printf("bottomData => %+v", bottomData)
		if err != nil {
			log.Printf("err => %s", err)
		}
	}
	dao.UpdateTableData(db, tableName, bottomData)
	ctx.JSON(http.StatusOK, res.Success(bottomData))
}
