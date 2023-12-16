package handle

import (
	"github.com/gin-gonic/gin"
	"log"
	"mongodb-first/dao"
	"mongodb-first/service"
	"net/http"
)

func (*StoreHandle) GetDatabaseColumnNameList(ctx *gin.Context) {
	//传入连接名
	article := service.GetColumnNameList(dao.Article)
	log.Printf("article => %+v", article)
	ctx.JSON(http.StatusOK, res.Success(article))
}
