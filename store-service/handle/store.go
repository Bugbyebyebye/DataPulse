package handle

import (
	mongo1 "commons/api/bottom/mongodb_first/gen"
	mysql1 "commons/api/bottom/mysql-first/gen"
	mysql2 "commons/api/bottom/mysql-second/gen"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"store-service/client"
)

// GetMysqlFirstData 获取Mysql1的数据
func (*StoreHandle) GetMysqlFirstData(ctx *gin.Context) {
	req := Req{Message: "这是服务端，mysql1你好", Target: "databaseList"}
	param, err := json.Marshal(req)
	if err != nil {
		log.Printf("err => %s", err)
	}

	var clientData Res
	var data []Res

	clientRes, err := client.MysqlFirstClient.GetMysqlFirstData(ctx, &mysql1.MysqlFirstReq{Param: param})
	if err != nil {
		log.Printf("mysql1 client err => %s", err)
	}
	err = json.Unmarshal(clientRes.Data, &clientData)
	if err != nil {
		log.Printf("err => %s", err)
	}
	data = append(data, clientData)

	log.Printf("bottom => %+v", data)
	ctx.JSON(200, res.Success(data))
}

func (*StoreHandle) GetMysqlSecondData(ctx *gin.Context) {
	req := Req{Message: "这是服务端，mysql2你好"}
	param, err := json.Marshal(req)
	if err != nil {
		log.Printf("err => %s", err)
	}
	clientRes, err := client.MysqlSecondClient.GetMysqlSecondData(ctx, &mysql2.MysqlSecondReq{Param: param})
	if err != nil {
		log.Printf("mysql2 client err => %s", err)
	}
	var data Res
	err = json.Unmarshal(clientRes.Data, &data)
	if err != nil {
		log.Printf("err => %s", err)
	}

	log.Printf("mysql2 bottom => %+v", data)
	ctx.JSON(200, res.Success(data))
}

func (*StoreHandle) GetMongoDbFirstData(ctx *gin.Context) {
	req := Req{Message: "这是服务端，Mongo1你好"}
	param, err := json.Marshal(req)
	if err != nil {
		log.Printf("err => %s", err)
	}
	clientRes, err := client.MongoDbFirstClient.GetMongoDbFirstData(ctx, &mongo1.MongoFirstReq{Param: param})
	if err != nil {
		log.Printf("mongo1 client err => %s", err)
	}
	var data Res
	err = json.Unmarshal(clientRes.Data, &data)
	if err != nil {
		log.Printf("err => %s", err)
	}

	log.Printf("mongo1 bottom => %+v", data)
	ctx.JSON(200, res.Success(data))
}
