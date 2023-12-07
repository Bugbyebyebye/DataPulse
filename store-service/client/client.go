package client

import (
	mongodb1 "commons/api/bottom/mongodb_first/gen"
	mysql1 "commons/api/bottom/mysql-first/gen"
	mysql2 "commons/api/bottom/mysql-second/gen"
	"commons/config"
	"commons/config/etcd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
)

var MysqlFirstClient mysql1.MysqlFirstServiceClient
var MysqlSecondClient mysql2.MysqlSecondServiceClient
var MongoDbFirstClient mongodb1.MongoDbFirstServiceClient

func InitDataBaseClient() {
	etcdRegister := etcd.NewResolver(config.Conf.ETCD.Addrs)
	resolver.Register(etcdRegister)

	mysqlFirst, err := grpc.Dial("etcd:///mysql1", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect mysql1 err => %s", err)
	}
	mysqlSecond, err := grpc.Dial("etcd:///mysql2", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect mysql2 err => %s", err)
	}
	mongoFirst, err := grpc.Dial("etcd:///mongo1", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect mongo1 err => %s", err)
	}

	MysqlFirstClient = mysql1.NewMysqlFirstServiceClient(mysqlFirst)
	MysqlSecondClient = mysql2.NewMysqlSecondServiceClient(mysqlSecond)
	MongoDbFirstClient = mongodb1.NewMongoDbFirstServiceClient(mongoFirst)
}
