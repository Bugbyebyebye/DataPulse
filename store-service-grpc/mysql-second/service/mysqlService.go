package mysql_second_service

import (
	mysql "commons/api/bottom/mysql-second/gen"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MysqlSecondService struct {
	mysql.UnimplementedMysqlSecondServiceServer
}

func (MysqlSecondService) GetMysqlSecondData(context.Context, *mysql.Req) (*mysql.Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMysqlSecondData not implemented")
}
