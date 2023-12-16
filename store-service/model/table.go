package model

import (
	"commons/util"
	"store-service/config"
)

type TableInfo struct {
	TableId      int    `json:"table_id" gorm:"column:table_id;primary_key;AUTO_INCREMENT"`
	Name         string `json:"table_name" gorm:"column:table_name"`
	DatabaseName string `json:"database_name" gorm:"column:database_name"`
	Type         int    `json:"table_type" gorm:"column:table_type"`
	FieldNum     int    `json:"field_num" gorm:"column:field_num"`
	State        int    `json:"state" gorm:"column:state"`
	CreateTime   int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime   int64  `json:"update_time" gorm:"column:update_time"`
}

func (TableInfo) TableName() string {
	return "t_table_info"
}

// InitTable 初始化数据表信息
func InitTable(tableName string, databaseName string, tableType int, fieldNum int) (int, error) {
	tableInfo := TableInfo{
		Name:         tableName,
		DatabaseName: databaseName,
		Type:         tableType,
		FieldNum:     fieldNum,
		State:        1,
		CreateTime:   util.GetUnixTime(),
		UpdateTime:   0}

	columns := []string{
		"table_name",
		"database_name",
		"table_type",
		"field_num",
		"state",
		"create_time",
	}
	err := config.System.Select(columns).Create(&tableInfo).Error
	return tableInfo.TableId, err
}

// GetTableInfo 获取数据库信息
func GetTableInfo(tableId int) (map[string]interface{}, error) {
	var tableInfo TableInfo
	var result map[string]interface{}
	err := config.System.Model(tableInfo).Where("table_id = ?", tableId).Find(&result).Error
	return result, err
}
