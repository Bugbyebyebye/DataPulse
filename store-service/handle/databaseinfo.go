package handle

import (
	"github.com/carlmjohnson/requests"
	"github.com/gin-gonic/gin"
	"log"
)

// GetDatabaseColumnNameList 获取数据库字段列表
func (*StoreHandle) GetDatabaseColumnNameList(ctx *gin.Context) {

	//数据源
	var dataFromList []DataSource
	var dataFrom DataSource

	var first []Database
	err := requests.URL("http://localhost:8085").
		Path("/getInfo").
		ToJSON(&first).Fetch(ctx)
	if err != nil {
		log.Printf("err => %+v", err)
	}
	var second []Database
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
