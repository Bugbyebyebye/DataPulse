package handle

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"store-service/client"
	"store-service/dao"
	//"store-service/model"
	mysql1 "commons/api/bottom/mysql-first/gen"
)

// CreateTable 创建数据仓库数据表
func (*StoreHandle) CreateTable(ctx *gin.Context) {
	var clientData Res
	var data []Res
	//id := ctx.GetInt("id")

	var jsonParam DataSource
	ctx.BindJSON(&jsonParam)
	//log.Printf("param %+v", jsonParam)
	println(jsonParam.SaveName)
	println(jsonParam.FromName)
	req := Req{Message: "这是服务端，mysql1你好", Target: "getColumnData", Param: jsonParam.DatabaseList}
	param, _ := json.Marshal(req)

	//TODO 检查是否有重名表
	//table := dao.Warehouse.HasTable(param.SaveName)
	//if table {
	//	ctx.JSON(http.StatusOK, res.Fail(4001, "表名已经存在"))
	//}
	//columns := jsonParam.DatabaseList[0].TableList[0].ColumnList

	//创建新数据库
	//dao.CreateTableBySQL(param.SaveName, columns)
	////创建用户数据库关联
	//tableId, err := model.InitTable(param.SaveName, len(columns))
	//if err != nil {
	//	log.Printf("err => %s", err)
	//}
	//err = model.InitTableUser(tableId, id)
	//if err != nil {
	//	log.Printf("err => %s", err)
	//}

	clientRes, err := client.MysqlFirstClient.GetMysqlFirstData(ctx, &mysql1.MysqlFirstReq{Param: param})
	if err != nil {
		log.Printf("mysql1 client err => %s", err)
	}
	json.Unmarshal(clientRes.Data, &clientData)
	data = append(data, clientData)

	log.Printf("bottom => %+v", data)
	//log.Printf("DatabaseList %+v", param.DatabaseList[0].TableList[0].ColumnList)
	ctx.JSON(http.StatusOK, res.Success(200))
}

// AlertTable 向数据仓库数据表中追加数据
func (*StoreHandle) AlertTable(ctx *gin.Context) {
	var param DataSource
	ctx.BindJSON(&param)

	log.Printf("param %+v", param)
	println(param.SaveName)
	println(param.FromName)

	columns := param.DatabaseList[0].TableList[0].ColumnList

	dao.AlertTableBySQL(param.SaveName, columns)

	ctx.JSON(http.StatusOK, res.Success(200))
}
