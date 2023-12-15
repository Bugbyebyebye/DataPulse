package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var Article *mongo.Client

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 建立连接
	conn, err := mongo.Connect(ctx,
		options.Client().
			// 连接地址
			ApplyURI("mongodb://222.186.50.126:20226").
			// 设置验证参数
			SetAuth(
				options.Credential{
					// 用户名
					Username: "root",
					// 密码
					Password: "maojiukeai1412",
				}).
			// 设置连接数
			SetMaxPoolSize(20))
	if err != nil {
		log.Println(err)
		return
	}
	// 测试连接
	err = conn.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("connect success!!!")
	Article = conn
}
