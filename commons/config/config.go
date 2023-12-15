package config

//配置 读取yml配置文件

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	Conf    = InitConfig()
	EnvConf = InitEnvConfig()
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
	c := &Config{viper: viper.New()}
	InitEnvConfig()
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

// InitEnvConfig 初始化Env
func InitEnvConfig() *Config {
	cs := &Config{}

	// 尝试从 .env 文件加载环境变量
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, loading environment variables directly.")
	}

	cs.ReaderServerConfigEnv()
	return cs
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
	//读取Smtp配置
	sm := &SmtpConfig{}
	sm.Host = c.viper.GetString("smtp.Host")
	sm.Username = c.viper.GetString("smtp.Username")
	sm.Password = c.viper.GetString("smtp.Password")
	sm.Fromname = c.viper.GetString("smtp.Fromname")
	c.SMTP = sm
}
func (c *Config) ReaderServerConfigEnv() {
	c.WEBHOOK = &SendWebHook{
		SendUrl: os.Getenv("SendUrl"),
	}
	c.MC = &MysqlConfig{
		Host:     os.Getenv("mysqlHost"),
		Name:     os.Getenv("mysqlName"),
		Password: os.Getenv("mysqlPassword"),
	}
}
