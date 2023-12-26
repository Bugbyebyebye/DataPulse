package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

//store-service 服务自己的配置
//数据库配置等

var (
	System     *gorm.DB
	Warehouse  *gorm.DB
	Warehouse2 *gorm.DB
	Conf       = InitConfig()
	err        error
)

type Config struct {
	SEVERURL *SeverUrlConfig
}

// SeverUrlConfig 服务地址配置
type SeverUrlConfig struct {
	AuthUrl  string
	LogsUrl  string
	TaskUrl  string
	StoreUrl string
	MYSQLF   string
	MYSQLS   string
	MongoDB  string
}

func init() {
	//系统数据库
	System, err = gorm.Open(mysql.Open("root:maojiukeai1412@tcp(222.186.50.126:20134)/df_system?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Printf("mysql error => %s", err)
	}
	if System.Error != nil {
		log.Printf("System error => %s", System.Error)
	}

	//数据仓库 数据库
	Warehouse, err = gorm.Open(mysql.Open("root:maojiukeai1412@tcp(222.186.50.126:20134)/df_warehouse?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Printf("mysql error => %s", err)
	}
	if Warehouse.Error != nil {
		log.Printf("Warehouse error => %s", Warehouse.Error)
	}

	Warehouse2, err = gorm.Open(mysql.Open("root:maojiukeai1412@tcp(222.186.50.126:20134)/df_warehouse2?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Printf("mysql error => %s", err)
	}
	if Warehouse2.Error != nil {
		log.Printf("Warehouse2 error => %s", Warehouse2.Error)
	}
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
	c.SEVERURL = &SeverUrlConfig{
		MYSQLF:  os.Getenv("MYSQLF"),
		MYSQLS:  os.Getenv("MYSQLS"),
		MongoDB: os.Getenv("MONGODB"),
	}
}

// GetDbByDatabaseName 根据数据库名 获取数据库操作指针
func GetDbByDatabaseName(databaseName string) *gorm.DB {
	if databaseName == "df_warehouse" {
		return Warehouse
	} else if databaseName == "df_warehouse2" {
		return Warehouse2
	}
	return nil
}

// QueryColumnData 传入字段值获取字段数据
func QueryColumnData(db *gorm.DB, tableName string, columnList []string) []map[string]interface{} {
	var data []map[string]interface{}
	err := db.Table(tableName).Select(columnList).Find(&data).Error
	if err != nil {
		log.Printf("mysql error => %s", err)
	}
	return data
}
