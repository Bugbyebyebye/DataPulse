package handle

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"mongodb-first/dao"
	"net/http"
	"time"
)

type SchoolWall struct {
	Id      string    `json:"id" bson:"_id"`
	Title   string    `json:"title" bson:"title"`
	Author  string    `json:"author" bson:"author"`
	Time    time.Time `json:"time" bson:"time"`
	Content string    `json:"content" bson:"content"`
	Url     string    `json:"url" bson:"url"`
}

func (*StoreHandle) GetDatabaseColumnNameList(ctx *gin.Context) {
	collection := dao.Article.Database("article").Collection("t_school_wall")
	find, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println(err)
		return
	}
	// 遍历查询结果
	var list []SchoolWall
	for find.Next(context.Background()) {
		var s SchoolWall
		// 解码绑定数据
		err = find.Decode(&s)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("%v", s)
		list = append(list, s)
	}

	ctx.JSON(http.StatusOK, res.Success(list))
}
