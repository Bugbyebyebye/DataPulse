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
	"strconv"
)

//获取数据库信息

// GetBottomDatabaseNameList 获取数据库字段列表
func (*StoreHandle) GetBottomDatabaseNameList(ctx *gin.Context) {

	//数据源
	var dataFromList []common.DataSource
	var dataFrom common.DataSource

	var first []common.Database
	err := requests.URL("http://localhost:8085").
		Path("/getInfo").
		ToJSON(&first).Fetch(ctx)
	if err != nil {
		log.Printf("err => %+v", err)
	}
	var second []common.Database
	err = requests.URL("http://localhost:8086").
		Path("/getInfo").
		ToJSON(&second).Fetch(ctx)
	if err != nil {
		log.Printf("err => %+v", err)
	}

	log.Printf("data => %+v", second)

	dataFrom.FromName = "mysql1"
	dataFrom.Databases = first
	dataFromList = append(dataFromList, dataFrom)

	dataFrom.FromName = "mysql2"
	dataFrom.Databases = second
	dataFromList = append(dataFromList, dataFrom)

	ctx.JSON(200, res.Success(dataFromList))
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
} // GetUserTableData 获取用户创建数据库的数据
func (*StoreHandle) GetUserTableData(ctx *gin.Context) {
	//TODO 传入一个表名
	databaseName := ctx.PostForm("database_name")
	tableName := ctx.PostForm("table_name")
	//TODO 根据数据表名返回全部数据
	db := config.GetDbByDatabaseName(databaseName)
	result, err := dao.QueryTableData(db, tableName)
	if err != nil {
		log.Printf("err => %s", err)
	}
	log.Printf("result => %s ", result)

	ctx.JSON(http.StatusOK, res.Success(result))
}
