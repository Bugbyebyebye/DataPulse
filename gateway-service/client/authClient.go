package client

import (
	authgrpc "commons/api/auth/gen"
	"commons/config"
	"commons/config/etcd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
)

var AuthClient authgrpc.TokenServiceClient

func InitAuthClient() {
	etcdRegister := etcd.NewResolver(config.Conf.ETCD.Addrs)
	resolver.Register(etcdRegister)

	authConn, err := grpc.Dial("etcd:///auth", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: err => %s", err)
	}
	AuthClient = authgrpc.NewTokenServiceClient(authConn)
}
