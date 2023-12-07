package router

import (
	auth "auth-service/router/auth"
	authservice "auth-service/service"
	authgrpc "commons/api/auth/gen"
	"commons/config"
	"commons/config/etcd"
	routers "commons/router"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
)

func InitRouter(r *gin.Engine) {
	rg := routers.New()
	rg.Route(&auth.AuthRouter{}, r)
}

// grp配置结构体
type gRpcConfig struct {
	Name         string
	Addr         string
	RegisterFunc func(*grpc.Server)
}

// RegisterGrpc 注册grpc 方法
func RegisterGrpc() *grpc.Server {
	c := gRpcConfig{
		Name: config.Conf.GRPC.Name,
		Addr: config.Conf.GRPC.Addr,
		RegisterFunc: func(g *grpc.Server) {
			authgrpc.RegisterTokenServiceServer(g, &authservice.AuthService{})
		},
	}
	server := grpc.NewServer()
	c.RegisterFunc(server)
	listen, err := net.Listen("tcp", c.Addr)
	if err != nil {
		log.Println("cannot listen")
	} else {
		log.Printf("grpc服务【%s】已启动! 运行在 => %s", c.Name, c.Addr)
	}
	go func() {
		err = server.Serve(listen)
		if err != nil {
			log.Println("app started error", err)
			return
		}
	}()
	return server
}

// EtcdServerRegister 将服务注册到etcd
func EtcdServerRegister() {
	etcdRegister := etcd.NewResolver(config.Conf.ETCD.Addrs)
	resolver.Register(etcdRegister)

	//注册grpc服务
	info := etcd.Server{
		Name:    config.Conf.GRPC.Name,
		Addr:    config.Conf.GRPC.Addr,
		Version: config.Conf.GRPC.Version,
		Weight:  config.Conf.GRPC.Weight,
	}

	r := etcd.NewRegister(config.Conf.ETCD.Addrs)
	_, err := r.Register(info, 2)
	if err != nil {
		log.Fatalln(err)
	}
}
