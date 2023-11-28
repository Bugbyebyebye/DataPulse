package data_service

import (
	data "commons/api/bottom/mysql-first"
	"context"
	"log"
)

type DataService struct {
	data.UnimplementedDataServiceServer
}

func (*DataService) GetDataInfo(ctx context.Context, req *data.Req) (res *data.Res, err error) {
	log.Printf("有人访问")
	id := req.GetId()
	if id == 1 {
		return &data.Res{Data: "hhhhhhhh"}, nil
	} else if id == 2 {
		return &data.Res{
			Data: "22k2jkfjdkfl",
		}, nil
	}
	return nil, nil
}
