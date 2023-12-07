package router

import (
	"github.com/gin-gonic/gin"
	"store-service/handle"
)

type StoreRouter struct {
}

func (*StoreRouter) Route(r *gin.Engine) {
	store := handle.New()

	r.GET("/getMysqlData", store.GetMysqlFirstData)
	r.GET("/getMysql2Data", store.GetMysqlSecondData)
	r.GET("/getMongoDbData", store.GetMongoDbFirstData)
}
