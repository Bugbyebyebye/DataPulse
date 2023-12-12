package mysql_first_service

import (
	mysql "commons/api/bottom/mysql-first/gen"
	"context"
	"encoding/json"
	"log"
	"mysql-first/handle"
)

type MysqlFirstService struct {
	mysql.UnimplementedMysqlFirstServiceServer
}

func (MysqlFirstService) GetMysqlFirstData(ctx context.Context, req *mysql.MysqlFirstReq) (res *mysql.MysqlFirstRes, err error) {
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
	} else if cq.Target == "getColumnData" {
		log.Printf("cq.param => %+v", cq.Param)
		result = handle.GetColumnData(cq.Param.([]interface{}))
	}

	//统一返回
	sr := ServerRes{Data: result}
	data, err := json.Marshal(sr)
	if err != nil {
		log.Printf("err => %s", err)
	}
	return &mysql.MysqlFirstRes{
		Data: data,
	}, nil
}
