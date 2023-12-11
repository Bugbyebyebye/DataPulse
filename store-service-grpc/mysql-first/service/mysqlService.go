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

		for _, database := range cq.Param.([]interface{}) {
			dbMap, ok := database.(map[string]interface{})
			if !ok {
				continue
			}
			dbName := dbMap["database_name"]
			tableList := dbMap["table_list"].([]interface{})
			for _, table := range tableList {
				tbMap, ok := table.(map[string]interface{})
				if !ok {
					continue
				}
				tbName := tbMap["table_name"]
				columnList := tbMap["column_list"].([]interface{})
				log.Printf("databaseName => %s tableName => %s columnlist => %s", dbName, tbName, columnList)

			}
		}
	}
	//统一返回
	sr := ServerRes{Message: "你好，我是 mysql1 grpc服务", Data: result}
	data, err := json.Marshal(sr)
	if err != nil {
		log.Printf("err => %s", err)
	}
	return &mysql.MysqlFirstRes{
		Data: data,
	}, nil
}
