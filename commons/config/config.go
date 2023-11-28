package config

//配置 读取yml配置文件

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var Conf = InitConfig()

type Config struct {
	viper *viper.Viper
	SC    *ServerConfig
	RC    *RedisConfig
	MC    *MysqlConfig
	GRPC  *GrpcConfig
	ETCD  *EtcdConfig
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Name string
	Addr string
}

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

// InitConfig 获取yml配置初始化
func InitConfig() *Config {
	c := &Config{viper: viper.New()}
	workDir, _ := os.Getwd()
	c.viper.SetConfigName("app")
	c.viper.SetConfigType("yml")
	c.viper.AddConfigPath(workDir + "/config")
	err := c.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	c.ReaderServerConfig()
	return c
}

// ReaderServerConfig 读取配置方法
func (c *Config) ReaderServerConfig() {
	//读取Gin Web服务配置
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.Name")
	sc.Addr = c.viper.GetString("server.Addr")
	c.SC = sc
	//读取GRPC 服务配置
	grpc := &GrpcConfig{}
	grpc.Name = c.viper.GetString("grpc.Name")
	grpc.Addr = c.viper.GetString("grpc.Addr")
	grpc.Version = c.viper.GetString("grpc.Version")
	grpc.Weight = c.viper.GetInt64("grpc.Weight")
	c.GRPC = grpc
	//读取ETCD 服务配置
	etcd := &EtcdConfig{}
	etcd.Name = c.viper.GetString("etcd.Name")
	var addrs []string
	err := c.viper.UnmarshalKey("etcd.Addrs", &addrs)
	if err != nil {
		log.Fatalln(err)
	}
	etcd.Addrs = addrs
	c.ETCD = etcd
	//读取Redis 配置
	rc := &RedisConfig{}
	rc.Host = c.viper.GetString("redis.Host")
	rc.Password = c.viper.GetString("redis.Password")
	rc.Db = c.viper.GetInt("redis.Db")
	c.RC = rc
	// 读取Mysql 服务配置
	mc := &MysqlConfig{}
	mc.Host = c.viper.GetString("mysql.Host")
	mc.Name = c.viper.GetString("mysql.Name")
	mc.Password = c.viper.GetString("mysql.Password")
	c.MC = mc
}
