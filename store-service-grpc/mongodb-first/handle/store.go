package handle

import (
	"context"
	"log"
	"mongodb-first/common"
	"mongodb-first/dao"
	"mongodb-first/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// GetDatabaseColumnNameList 获取数据表信息
func (*StoreHandle) GetDatabaseColumnNameList(ctx *gin.Context) {
	//传入连接名
	article := service.GetColumnNameList(dao.Article)
	//log.Printf("article => %+v", article)
	ctx.JSON(http.StatusOK, article)
}

// GetColumnData 获取指定字段数据
func (*StoreHandle) GetColumnData(ctx *gin.Context) {
	var table common.Table

	err := ctx.BindJSON(&table)
	if err != nil {
		log.Printf("err => %s", err)
	}
	result := service.GetColumnData(table)
	//log.Printf("result => %+v", result)

	ctx.JSON(http.StatusOK, result)
}

type SchoolWall struct {
	Id      string    `json:"id" bson:"_id"`
	Title   string    `json:"title" bson:"title"`
	Author  string    `json:"author" bson:"author"`
	Time    time.Time `json:"time" bson:"time"`
	Content string    `json:"content" bson:"content"`
	Url     string    `json:"url" bson:"url"`
}

// GetSchoolWallCleanedInfo 返回校园墙数据表中非空数据，附带一个脏数据数量，但是不知道怎么把这个数一起返回，所以目前只打印了出来
func (*StoreHandle) GetSchoolWallCleanedInfo(ctx *gin.Context) {
	collection := dao.Article.Database("article").Collection("t_school_wall")
	find, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println(err)
		return
	}

	// 遍历查询结果
	var list []SchoolWall
	var count int
	for find.Next(context.Background()) {
		var s SchoolWall
		// 解码绑定数据
		err = find.Decode(&s)
		if err != nil {
			log.Println(err)
			return
		}

		if s.Title != "" && s.Author != "" && s.Content != "" && s.Url != "" {
			list = append(list, s)
		} else {
			count++
		}
	}
	log.Println(count)
	ctx.JSON(http.StatusOK, res.Success(list))
}

type WeChat struct {
	Id      string `json:"id" bson:"_id"`
	Title   string `json:"title" bson:"title"`
	Author  string `json:"author" bson:"author"`
	Time    string `json:"time" bson:"time"`
	Content string `json:"content" bson:"content"`
	Url     string `json:"url" bson:"url"`
}

// GetWeChatCleanedInfo 返回WeChat数据表中非空数据，附带打印一个脏数据数量
func (*StoreHandle) GetWeChatCleanedInfo(ctx *gin.Context) {
	collection := dao.Article.Database("article").Collection("t_wechat")
	find, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println(err)
		return
	}

	// 遍历查询结果
	var list []WeChat
	var count int
	for find.Next(context.Background()) {
		var s WeChat
		// 解码绑定数据
		err = find.Decode(&s)
		if err != nil {
			log.Println(err)
			return
		}

		if s.Title != "" && s.Author != "" && s.Content != "" && s.Url != "" {
			list = append(list, s)
		} else {
			count++
		}
	}
	log.Println(count)
	ctx.JSON(http.StatusOK, res.Success(list))
}
