package model

import (
	"commons/util"
	"store-service/config"
)

type DatabaseInfo struct {
	DatabaseId   int    `json:"database_id" gorm:"column:database_id;primary_key;AUTO_INCREMENT"`
	DatabaseName string `json:"database_name" gorm:"column:database_name"`
	TableNum     int    `json:"table_num" gorm:"column:table_num"`
	State        int    `json:"state" gorm:"column:state"`
	CreateTime   int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime   int64  `json:"update_time" gorm:"column:update_time"`
}

func (DatabaseInfo) TableName() string {
	return "t_database_info"
}

// GetDatabaseIdByName 根据数据库名称获取数据库id
func GetDatabaseIdByName(databaseName string) (int, error) {
	var databaseId int
	err := config.System.Model(DatabaseInfo{}).Select("database_id").Where("database_name = ?", databaseName).Find(&databaseId).Error
	return databaseId, err
}

func GetTableNumByName(databaseName string) (int, error) {
	var tableNum int
	err := config.System.Model(DatabaseInfo{}).Select("table_num").Where("database_name = ?", databaseName).Find(&tableNum).Error
	return tableNum, err
}

// AddTableNumByName 新增表数量
func AddTableNumByName(databaseName string) error {
	tableNum, _ := GetTableNumByName(databaseName)
	err := config.System.Model(DatabaseInfo{}).Updates(map[string]interface{}{
		"table_num":   tableNum + 1,
		"update_time": util.GetUnixTime(),
	}).Where("database_name = ?", databaseName).Error
	return err
}

// SubTableNumByName 删减表数量
func SubTableNumByName(databaseName string) error {
	tableNum, _ := GetTableNumByName(databaseName)
	err := config.System.Model(DatabaseInfo{}).Updates(map[string]interface{}{
		"table_num":   tableNum - 1,
		"update_time": util.GetUnixTime(),
	}).Where("database_name = ?", databaseName).Error
	return err
}
