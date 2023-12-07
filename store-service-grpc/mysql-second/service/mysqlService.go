package mysql_second_service

import (
	mysql "commons/api/bottom/mysql-second/gen"
	"context"
	"encoding/json"
	"log"
)

type MysqlSecondService struct {
	mysql.UnimplementedMysqlSecondServiceServer
}

type ClientReq struct {
	Message string `json:"message"`
}

type ServerRes struct {
	Message string `json:"message"`
}

func (MysqlSecondService) GetMysqlSecondData(ctx context.Context, req *mysql.MysqlSecondReq) (res *mysql.MysqlSecondRes, err error) {
	var cq ClientReq
	err = json.Unmarshal(req.Param, &cq)
	log.Printf("来自客户端的请求信息为 =>  %+v", cq)

	sr := ServerRes{Message: "你好，我是 mysql2 grpc服务"}
	data, err := json.Marshal(sr)
	if err != nil {
		log.Printf("err => %s", err)
	}
	return &mysql.MysqlSecondRes{
		Data: data,
	}, nil
}
