package config

//配置 读取yml配置文件

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

var (
	Conf = InitConfig()
)

type Config struct {
	viper   *viper.Viper
	SC      *ServerConfig
	RC      *RedisConfig
	MC      *MysqlConfig
	GRPC    *GrpcConfig
	ETCD    *EtcdConfig
	SMTP    *SmtpConfig
	WEBHOOK *SendWebHook
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Name string
	Addr string
}

// MysqlConfig MySQL配置
type MysqlConfig struct {
	Host     string
	Name     string
	Password string
}

// RedisConfig Redis 配置
type RedisConfig struct {
	Host     string
	Password string
	Db       int
}

type SmtpConfig struct {
	Host     string
	Username string
	Password string
	Fromname string
}

// GrpcConfig grpc配置
type GrpcConfig struct {
	Name    string
	Addr    string
	Version string
	Weight  int64
}

// EtcdConfig etcd 配置
type EtcdConfig struct {
	Name  string
	Addrs []string
}

// SendWebHook SendWebHookUrl SendWebHook 发送日志通知
type SendWebHook struct {
	SendUrl string
}

// InitConfig 获取yml配置初始化
func InitConfig() *Config {
	c := &Config{}

	// 尝试从 .env 文件加载环境变量
	if err := godotenv.Load(); err != nil {
		fmt.Println("没有.env文件，尝试从系统环境变量中获取")
	}

	c.ReaderServerConfigEnv()
	return c
}

func (c *Config) ReaderServerConfigEnv() {
	RedisDb := os.Getenv("redisDb")       // 获取环境变量的值
	RedisDB, err := strconv.Atoi(RedisDb) // 将字符串转换为整数
	if err != nil {
		// 处理转换错误
		fmt.Println(err)
		fmt.Println("类型转换错误")
	}
	//GrpcWeigit := os.Getenv("GRPCWEIGHT")
	//GRPCWeight, err := strconv.Atoi(GrpcWeigit)
	//if err != nil {
	//	fmt.Println("类型转换错误")
	//}
	//GRPCWeightInt64 := int64(GRPCWeight) // 将 int 转换为 int64
	////转换为切片
	//etcdAddr := os.Getenv("ETCDADDR")    // 获取环境变量的值
	//addresses := []string{}
	//
	//if etcdAddr != "" {
	//	addresses = strings.Split(etcdAddr, ",")
	//	for i, addr := range addresses {
	//		addresses[i] = strings.TrimSpace(addr) // 去除地址两端的空格
	//	}
	//}
	c.WEBHOOK = &SendWebHook{
		SendUrl: os.Getenv("SendUrl"),
	}
	c.SMTP = &SmtpConfig{
		Host:     os.Getenv("smtpHost"),
		Username: os.Getenv("smtpUsername"),
		Password: os.Getenv("smtpPassword"),
		Fromname: os.Getenv("smtpFromname"),
	}
	c.MC = &MysqlConfig{
		Host:     os.Getenv("MYSQLHOST"),
		Name:     os.Getenv("MYSQLNAME"),
		Password: os.Getenv("MYSQLPASSWORD"),
	}
	c.RC = &RedisConfig{
		Host:     os.Getenv("redisHost"),
		Password: os.Getenv("redisPassword"),
		Db:       RedisDB,
	}
	c.SC = &ServerConfig{
		Name: os.Getenv("SERVERNAME"),
		Addr: os.Getenv("SERVERADDR"),
	}
	//c.GRPC = &GrpcConfig{
	//	Name:    os.Getenv("GRPCNAME"),
	//	Addr:    os.Getenv("GRPCADDR"),
	//	Version: os.Getenv("GRPCVERSION"),
	//	Weight:  GRPCWeightInt64,
	//}
	//c.ETCD = &EtcdConfig{
	//	Name:  os.Getenv("ETCDNAME"),
	//	Addrs: addresses,
	//}
}
