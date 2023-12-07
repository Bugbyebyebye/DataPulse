package handle

import (
	mysql1 "commons/api/bottom/mysql-first/gen"
	"commons/result"
	"log"

	"github.com/gin-gonic/gin"
	"store-service/client"
)

type StoreHandle struct {
}

func New() *StoreHandle {
	return &StoreHandle{}
}

var res result.Result

type MysqlFirstData struct {
}

func (*StoreHandle) GetMysqlFirstData(ctx *gin.Context) {
	var param mysql1.Req
	data, err := client.MysqlFirstClient.GetMysqlFirstData(ctx, &mysql1.Req{Param: param})
	if err != nil {
		log.Printf("mysql1 client err => %s", err)
	}

	log.Printf("data => %+v", data)
	ctx.JSON(200, res.Success())
}

func (*StoreHandle) GetMysqlSecondData(ctx *gin.Context) {

}

func (*StoreHandle) GetMongoDbFirstData(ctx *gin.Context) {

}
