package mongodb_service

import (
	mongo "commons/api/bottom/mongodb_first/gen"
	"context"
	"encoding/json"
	"log"
)

type MongoFirstService struct {
	mongo.UnimplementedMongoDbFirstServiceServer
}

type ClientReq struct {
	Message string `json:"message"`
}

type ServerRes struct {
	Message string `json:"message"`
}

func (MongoFirstService) GetMongoDbFirstData(ctx context.Context, req *mongo.MongoFirstReq) (res *mongo.MongoFirstRes, err error) {
	var cq ClientReq
	err = json.Unmarshal(req.Param, &cq)
	log.Printf("来自客户端的请求信息为 =>  %+v", cq)

	sr := ServerRes{Message: "你好，我是 mongo1 grpc服务"}
	data, err := json.Marshal(sr)
	if err != nil {
		log.Printf("err => %s", err)
	}
	return &mongo.MongoFirstRes{
		Data: data,
	}, nil
}
