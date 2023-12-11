package mysql_second_service

import (
	mysql "commons/api/bottom/mysql-second/gen"
	"context"
	"encoding/json"
	"log"
	"mysql-second/handle"
)

type MysqlSecondService struct {
	mysql.UnimplementedMysqlSecondServiceServer
}

type ClientReq struct {
	Message string `json:"message"`
	Target  string `json:"target"`
}

type ServerRes struct {
	Message string      `json:"message"`
	Data    interface{} `json:"bottom"`
}

func (MysqlSecondService) GetMysqlSecondData(ctx context.Context, req *mysql.MysqlSecondReq) (res *mysql.MysqlSecondRes, err error) {
	var result interface{}

	//统一解析请求
	var cq ClientReq
	err = json.Unmarshal(req.Param, &cq)
	log.Printf("来自客户端的请求信息为 =>  %+v", cq)

	//数据操作
	//1.获取数据库全部表名和字段名
	if cq.Target == "databaseList" {
		result = handle.GetColumnNameList()
		//log.Printf("list1 => %+v", result)
	}

	sr := ServerRes{Message: "你好，我是 mysql2 grpc服务", Data: result}
	data, err := json.Marshal(sr)
	if err != nil {
		log.Printf("err => %s", err)
	}
	return &mysql.MysqlSecondRes{
		Data: data,
	}, nil
}
