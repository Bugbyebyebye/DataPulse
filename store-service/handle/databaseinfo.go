package handle

import (
	mysql1 "commons/api/bottom/mysql-first/gen"
	mysql2 "commons/api/bottom/mysql-second/gen"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"store-service/client"
)

// GetDatabaseColumnNameList 获取数据库字段列表
func (*StoreHandle) GetDatabaseColumnNameList(ctx *gin.Context) {
	req := Req{Message: "获取数据库字段列表", Target: "databaseList"}
	param, err := json.Marshal(req)
	if err != nil {
		log.Printf("err => %s", err)
	}

	//数据源
	var dataFromList []DataSource
	var dataFrom DataSource
	var clientData Res

	clientRes, err := client.MysqlFirstClient.GetMysqlFirstData(ctx, &mysql1.MysqlFirstReq{Param: param})
	if err != nil {
		log.Printf("mysql1 client err => %s", err)
	}
	json.Unmarshal(clientRes.Data, &clientData)
	dataFrom.FromName = "mysql1"
	dataFrom.Databases = clientData.Data
	log.Printf("clientData.Data => %+v", clientData.Data)
	dataFromList = append(dataFromList, dataFrom)

	clientRes2, err := client.MysqlSecondClient.GetMysqlSecondData(ctx, &mysql2.MysqlSecondReq{Param: param})
	if err != nil {
		log.Printf("mysql2 client err => %s", err)
	}
	json.Unmarshal(clientRes2.Data, &clientData)
	dataFrom.FromName = "mysql2"
	dataFrom.Databases = clientData.Data
	dataFromList = append(dataFromList, dataFrom)

	ctx.JSON(200, res.Success(dataFromList))
}
