package handle

import (
	"github.com/carlmjohnson/requests"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"store-service/common"
	"store-service/config"
	"store-service/dao"
	"store-service/model"
	"store-service/service"
	"strconv"
)

//获取数据库信息

// GetBottomDatabaseNameList 获取数据库字段列表
func (*StoreHandle) GetBottomDatabaseNameList(ctx *gin.Context) {
	//数据源
	var result []common.Table
	role := ctx.Request.Header.Get("role")

	if role == "admin" {
		//登录用户为管理员 从底层数据库拉取数据源供管理员使用
		result = service.GetBottomTableInfo(ctx)
	} else if role == "user" {
		//登录用户为普通用户 拉取数据仓库中的公共数据源供用户使用
		result = service.GetWarehousePublicTableInfo()
	}

	ctx.JSON(200, res.Success(result))
}

// GetUserDatabaseNameList 获取用户创建的数据表信息
func (*StoreHandle) GetUserDatabaseNameList(ctx *gin.Context) {
	//从token中拿到用户id
	idStr := ctx.GetHeader("id")
	id, _ := strconv.Atoi(idStr)

	//去用户数据库关联表拿到数据表信息
	list, err := model.GetTableIdList(id)
	if err != nil {
		log.Printf("err => %s", err)
	}
	log.Printf("list => %+v", list)
	//返回数据表相关信息
	var result []map[string]interface{}
	for _, v := range list {
		info, err := model.GetTableInfo(v)
		if err != nil {
			log.Printf("err => %s", err)
		}
		log.Printf("info => %+v", info)
		result = append(result, info)
	}

	ctx.JSON(http.StatusOK, res.Success(result))
}

// GetUserTableData 获取用户创建数据库的数据
func (*StoreHandle) GetUserTableData(ctx *gin.Context) {
	//传入一个表名
	databaseName := ctx.PostForm("database_name")
	tableName := ctx.PostForm("table_name")
	//根据数据表名返回全部数据
	db := config.GetDbByDatabaseName(databaseName)
	result, err := dao.QueryTableData(db, tableName)
	if err != nil {
		log.Printf("err => %s", err)
	}
	log.Printf("result => %s ", result)

	ctx.JSON(http.StatusOK, res.Success(result))
}

type TargetInfo struct {
	Database   string   `json:"database_name"`
	Table      string   `json:"table_name"`
	ColumnList []string `json:"column_list"`
}

// GetWarehouseDatabaseNameList 获取数据仓库中数据库数据表信息
func (*StoreHandle) GetWarehouseDatabaseNameList(ctx *gin.Context) {
	var targetList []common.Table
	databaseNames := []string{"df_warehouse", "df_warehouse2"}
	for _, database := range databaseNames {
		db := config.GetDbByDatabaseName(database)
		//log.Printf("db => %+v", db)
		tableNames := dao.GetAllTableName(db)
		for _, table := range tableNames {
			columnNames := dao.GetAllColumnName(table, db)
			var target common.Table
			target.SourceName = "mysql"
			target.DatabaseName = database
			target.TableName = table
			target.ColumnList = columnNames

			targetList = append(targetList, target)
		}
	}

	ctx.JSON(http.StatusOK, res.Success(targetList))
}

type IndexData struct {
	DatabaseNum     int `json:"database_num"`
	PublicTableNum  int `json:"public_table_num"`
	PrivateTableNum int `json:"private_table_num"`
	ApiTotalNum     int `json:"api_total_num"`
	ApiWarningNum   int `json:"api_warning_num"`
	ApiRunningNum   int `json:"api_running_num"`
}

// GetIndexTableData 获取首页需要的6个统计数据
func (*StoreHandle) GetIndexTableData(ctx *gin.Context) {
	idStr := ctx.Request.Header.Get("id")
	id, _ := strconv.Atoi(idStr)

	//获取数据表数据
	//公共表
	publicTableNum, err := model.GetPublicTableNum()
	if err != nil {
		log.Printf("err => %s", err)
	}
	//私有表
	var privateTableNum int
	list, err := model.GetTableIdList(id)
	if err != nil {
		log.Printf("err => %s", err)
	}
	for _, v := range list {
		num, err := model.GetPersonalTableNum(v)
		if err != nil {
			log.Printf("err => %s", err)
		}
		privateTableNum += int(num)
	}

	var result map[string]int
	//获取api统计数据
	err = requests.URL("http://task-service:8084").
		Path("/searchstate").
		BodyJSON(map[string]interface{}{
			"user_id": idStr,
		}).
		ToJSON(&result).Fetch(ctx)
	if err != nil {
		log.Printf("err => %s", err)
	}
	log.Printf("result => %+v", result)

	var data IndexData
	data.DatabaseNum = 2
	data.PublicTableNum = int(publicTableNum)
	data.PrivateTableNum = privateTableNum
	data.ApiTotalNum = result["1"]
	data.ApiWarningNum = result["3"]
	data.ApiRunningNum = result["2"]
	ctx.JSON(http.StatusOK, res.Success(data))
}
