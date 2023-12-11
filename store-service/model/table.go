package model

import (
	"commons/util"
	"store-service/dao"
)

type TableInfo struct {
	TableId    int    `json:"table_id" gorm:"column:table_id;primary_key;AUTO_INCREMENT"`
	Name       string `json:"table_name" gorm:"column:table_name"`
	FieldNum   int    `json:"field_num" gorm:"column:field_num"`
	State      int    `json:"state" gorm:"column:state"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time"`
}

func (TableInfo) TableName() string {
	return "t_table_info"
}

// InitTable 初始化数据表信息
func InitTable(tableName string, fieldNum int) (int, error) {
	tableInfo := TableInfo{
		Name:       tableName,
		FieldNum:   fieldNum,
		State:      1,
		CreateTime: util.GetUnixTime(),
		UpdateTime: 0}

	err := dao.System.Select("table_name", "field_num", "state", "create_time").Create(&tableInfo).Error
	return tableInfo.TableId, err
}
