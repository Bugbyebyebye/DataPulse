package register

//
////注册GRPC服务
//
//import (
//	data2 "commons/api/bottom/mysql-second"
//	"commons/config"
//	"commons/config/etcd"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/resolver"
//	"log"
//	DataService "mysql-second/service"
//	"net"
//)
//
//// grp配置结构体
//type gRpcConfig struct {
//	Name         string
//	Addr         string
//	RegisterFunc func(*grpc.Server)
//}
//
//// GrpcRegister 注册grpc 方法
//func GrpcRegister() {
//	c := gRpcConfig{
//		Name: config.Conf.GRPC.Name,
//		Addr: config.Conf.GRPC.Addr,
//		RegisterFunc: func(g *grpc.Server) {
//			data2.RegisterDataServiceServer(g, &DataService.DataService{})
//		},
//	}
//	server := grpc.NewServer()
//	c.RegisterFunc(server)
//	listen, err := net.Listen("tcp", c.Addr)
//	if err != nil {
//		log.Println("cannot listen")
//	} else {
//		log.Printf("grpc服务【%s】已启动! 运行在 => %s", c.Name, c.Addr)
//	}
//
//	err2 := server.Serve(listen)
//	if err2 != nil {
//		log.Println("server started error", err)
//		return
//	}
//}
//
//// EtcdServerRegister 将服务注册到etcd
//func EtcdServerRegister() {
//	etcdRegister := etcd.NewResolver(config.Conf.ETCD.Addrs)
//	resolver.Register(etcdRegister)
//
//	//注册grpc服务
//	info := etcd.Server{
//		Name:    config.Conf.GRPC.Name,
//		Addr:    config.Conf.GRPC.Addr,
//		Version: config.Conf.GRPC.Version,
//		Weight:  config.Conf.GRPC.Weight,
//	}
//
//	r := etcd.NewRegister(config.Conf.ETCD.Addrs)
//	_, err := r.Register(info, 2)
//	if err != nil {
//		log.Fatalln(err)
//	}
//}
