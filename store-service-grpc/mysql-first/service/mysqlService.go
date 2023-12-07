package mysql_first_service

import (
	mysql "commons/api/bottom/mysql-first/gen"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MysqlFirstService struct {
	mysql.UnimplementedMysqlFirstServiceServer
}

func (MysqlFirstService) GetMysqlFirstData(context.Context, *mysql.Req) (*mysql.Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMysqlFirstData not implemented")
}
