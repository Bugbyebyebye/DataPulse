package handle

import (
	mongo1 "commons/api/bottom/mongodb_first/gen"
	mysql1 "commons/api/bottom/mysql-first/gen"
	mysql2 "commons/api/bottom/mysql-second/gen"
	"commons/result"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"store-service/client"
)

var res result.Result

type MysqlFirstReq struct {
	Message string `json:"message"`
}
type MysqlFirstRes struct {
	Message string `json:"message"`
}

// GetMysqlFirstData 获取Mysql1的数据
func (*StoreHandle) GetMysqlFirstData(ctx *gin.Context) {
	req := MysqlFirstReq{Message: "这是服务端，mysql1你好"}
	param, err := json.Marshal(req)
	if err != nil {
		log.Printf("err => %s", err)
	}

	clientRes, err := client.MysqlFirstClient.GetMysqlFirstData(ctx, &mysql1.MysqlFirstReq{Param: param})
	if err != nil {
		log.Printf("mysql1 client err => %s", err)
	}

	var data MysqlFirstRes
	err = json.Unmarshal(clientRes.Data, &data)
	if err != nil {
		log.Printf("err => %s", err)
	}

	log.Printf("data => %+v", data)
	ctx.JSON(200, res.Success(data))
}

type MysqlSecondReq struct {
	Message string `json:"message"`
}
type MysqlSecondRes struct {
	Message string `json:"message"`
}

func (*StoreHandle) GetMysqlSecondData(ctx *gin.Context) {
	req := MysqlSecondReq{Message: "这是服务端，mysql2你好"}
	param, err := json.Marshal(req)
	if err != nil {
		log.Printf("err => %s", err)
	}
	clientRes, err := client.MysqlSecondClient.GetMysqlSecondData(ctx, &mysql2.MysqlSecondReq{Param: param})
	if err != nil {
		log.Printf("mysql2 client err => %s", err)
	}
	var data MysqlSecondRes
	err = json.Unmarshal(clientRes.Data, &data)
	if err != nil {
		log.Printf("err => %s", err)
	}

	log.Printf("mysql2 data => %+v", data)
	ctx.JSON(200, res.Success(data))
}

type MongoFirstReq struct {
	Message string `json:"message"`
}
type MongoFirstRes struct {
	Message string `json:"message"`
}

func (*StoreHandle) GetMongoDbFirstData(ctx *gin.Context) {
	req := MongoFirstReq{Message: "这是服务端，Mongo1你好"}
	param, err := json.Marshal(req)
	if err != nil {
		log.Printf("err => %s", err)
	}
	clientRes, err := client.MongoDbFirstClient.GetMongoDbFirstData(ctx, &mongo1.MongoFirstReq{Param: param})
	if err != nil {
		log.Printf("mongo1 client err => %s", err)
	}
	var data MongoFirstRes
	err = json.Unmarshal(clientRes.Data, &data)
	if err != nil {
		log.Printf("err => %s", err)
	}

	log.Printf("mongo1 data => %+v", data)
	ctx.JSON(200, res.Success(data))
}
