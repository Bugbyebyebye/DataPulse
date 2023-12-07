package mysql_first_service

import (
	mysql "commons/api/bottom/mysql-first/gen"
	"context"
	"encoding/json"
	"log"
)

type MysqlFirstService struct {
	mysql.UnimplementedMysqlFirstServiceServer
}

type ClientReq struct {
	Message string `json:"message"`
}

type ServerRes struct {
	Message string `json:"message"`
}

func (MysqlFirstService) GetMysqlFirstData(ctx context.Context, req *mysql.MysqlFirstReq) (res *mysql.MysqlFirstRes, err error) {
	var cq ClientReq
	err = json.Unmarshal(req.Param, &cq)
	log.Printf("来自客户端的请求信息为 =>  %+v", cq)

	sr := ServerRes{Message: "你好，我是 mysql1 grpc服务"}
	data, err := json.Marshal(sr)
	if err != nil {
		log.Printf("err => %s", err)
	}
	return &mysql.MysqlFirstRes{
		Data: data,
	}, nil
}
