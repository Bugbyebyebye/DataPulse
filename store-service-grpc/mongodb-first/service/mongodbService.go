package mongodb_service

import (
	mongo "commons/api/bottom/mongodb_first/gen"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MongoFirstService struct {
	mongo.UnimplementedMongoDbFirstServiceServer
}

func (MongoFirstService) GetMongoDbFirstData(context.Context, *mongo.Req) (*mongo.Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMongoDbFirstData not implemented")
}
