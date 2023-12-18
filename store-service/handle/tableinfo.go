package handle

import (
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
