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

	var req common.CreateReq
	err := ctx.BindJSON(&req)
	if err != nil {
		log.Printf("err => %s", err)
	}

	targetDatabase := req.TargetDatabase //目标数据库名
	targetTable := req.TargetTable       //目标表名
	targetType := req.TargetType         //目标表类型
	tableList := req.TableList           //数据表列表

	db := config.GetDbByDatabaseName(targetDatabase)

	//检查是否有重名表
	tableNameExist := dao.GetAllTableName(db)
	if util.In(targetTable, tableNameExist) {
		ctx.JSON(http.StatusOK, res.Fail(4001, "表名已存在，请重新选择"))
		return
	}

	//将字段列表去重
	columns := common.GetUniqueColumnList(tableList)
	log.Printf("columns => %+v", columns)

	//创建新数据库
	dao.CreateTableBySQL(db, targetTable, columns)
	//创建用户数据库关联
	tableId, err := model.InitTable(targetTable, targetDatabase, targetType, len(columns))
	if err != nil {
		log.Printf("err => %s", err)
	}
	//初始化数据表和用户关联
	err = model.InitTableUser(tableId, id)
	if err != nil {
		log.Printf("err => %s", err)
	}

	log.Printf("tableList => %+v", tableList)
	//根据数据源选择底层数据来源
	for i, t := range tableList {

		var bottom []map[string]interface{}
		if t.SourceName == "mysql" {
			bottom = dao.GetDataByColumnList(t)
			log.Printf("bottom => %+v\n", bottom)
		} else if t.SourceName == "mysql1" {
			err := requests.URL("http://mysql-first:8085").
				Path("/getColumnData").
				BodyJSON(&t).
				ToJSON(&bottom).
				Fetch(ctx)
			log.Printf("bottom => %+v\n", bottom)
			if err != nil {
				log.Printf("err => %s", err)
			}
		} else if t.SourceName == "mysql2" {
			err := requests.URL("http://mysql-second:8086").
				Path("/getColumnData").
				BodyJSON(&t).
				ToJSON(&bottom).
				Fetch(ctx)
			log.Printf("bottom => %+v", bottom)
			if err != nil {
				log.Printf("err => %s", err)
			}
		} else if t.SourceName == "mongodb1" {
			err := requests.URL("http://mongodb-first:8087").
				Path("/getColumnData").
				BodyJSON(&t).
				ToJSON(&bottom).
				Fetch(ctx)
			log.Printf("bottom => %+v", bottom)
			if err != nil {
				log.Printf("err => %s", err)
			}
		}

		if i == 0 {
			dao.InitTableData(db, targetTable, bottom)
		} else {
			dao.UpdateTableData(db, targetTable, t.RelateFlag, bottom)
		}
	}

	//根据底层数据 新增数据
	fmt.Println("Data inserted successfully.")

	ctx.JSON(http.StatusOK, res.Success("创建数据表成功"))
}

// AlertTable 向数据仓库数据表中追加字段
func (*StoreHandle) AlertTable(ctx *gin.Context) {
	var req common.CreateReq
	err := ctx.BindJSON(&req)
	if err != nil {
		log.Printf("err => %s", err)
	}

	targetDatabase := req.TargetDatabase //目标数据库名
	targetTable := req.TargetTable       //目标表名
	tableList := req.TableList           //数据表列表

	//匹配目标数据库
	db := config.GetDbByDatabaseName(targetDatabase)

	//获取去重后的字段列表
	columns := common.GetUniqueColumnList(tableList)

	//向数据表中追加新增字段
	dao.AlertTableBySQL(db, targetTable, columns)

	//根据数据源选择数据来源
	for _, t := range tableList {
		var bottom []map[string]interface{}
		if t.SourceName == "mysql" {
			bottom = dao.GetDataByColumnList(t)
			log.Printf("bottom => %+v\n", bottom)
		} else if t.SourceName == "mysql1" {
			err := requests.URL("http://mysql-first:8085").
				Path("/getColumnData").
				BodyJSON(&t).
				ToJSON(&bottom).
				Fetch(ctx)
			log.Printf("bottom => %+v\n", bottom)
			if err != nil {
				log.Printf("err => %s", err)
			}
		} else if t.SourceName == "mysql2" {
			err := requests.URL("http://mysql-second:8086").
				Path("/getColumnData").
				BodyJSON(&t).
				ToJSON(&bottom).
				Fetch(ctx)
			log.Printf("bottom => %+v\n", bottom)
			if err != nil {
				log.Printf("err => %s", err)
			}
		} else if t.SourceName == "mongodb1" {
			err := requests.URL("http://mongodb-first:8087").
				Path("/getColumnData").
				BodyJSON(&t).
				ToJSON(&bottom).
				Fetch(ctx)
			log.Printf("bottom => %+v\n", bottom)
			if err != nil {
				log.Printf("err => %s", err)
			}
		}

		//将数据组合
		dao.UpdateTableData(db, targetTable, t.RelateFlag, bottom)
	}

	ctx.JSON(http.StatusOK, res.Success("数据表追加字段成功"))
}
