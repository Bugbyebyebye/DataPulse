package data_service

import (
	data "commons/api/bottom/mysql-second"
	"context"
	"log"
)

type DataService struct {
	data.UnimplementedDataServiceServer
}

func (*DataService) GetDataInfo(ctx context.Context, req *data.Req) (res *data.Res, err error) {
	log.Printf("有人访问")

	return nil, nil
}
