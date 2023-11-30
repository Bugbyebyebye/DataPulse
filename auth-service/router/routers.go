package router

import (
	router "auth-service/router/auth"
	authservice "auth-service/service"
	authgrpc "commons/api/auth/gen"
	"commons/config"
	"commons/config/etcd"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
)

// Router 路由接口
type Router interface {
	Route(r *gin.Engine)
}

// RegisterRouter 注册路由初始化方法
type RegisterRouter struct {
}

func New() *RegisterRouter {
	return &RegisterRouter{}
}

// Route 路由封装类
func (*RegisterRouter) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

func InitRouter(r *gin.Engine) {
	rg := New()
	rg.Route(&router.AuthRouter{}, r)
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
