package model

import (
	"commons/util"
	"store-service/config"
)

type DatabaseTableRelate struct {
	DatabaseTableId int   `json:"database_table_id" gorm:"column:database_id;primary_key;AUTO_INCREMENT"`
	DatabaseId      int   `json:"database_id" gorm:"column:database_id"`
	TableId         int   `json:"table_id" gorm:"column:table_id"`
	State           int   `json:"state" gorm:"column:state"`
	CreateTime      int64 `json:"create_time" gorm:"column:create_time"`
	UpdateTime      int64 `json:"update_time" gorm:"column:update_time"`
}

func (DatabaseTableRelate) TableName() string {
	return "t_database_table_relate"
}

// InsertData 插入数据
func InsertData(databaseId int, tableId int) error {
	err := config.System.Model(DatabaseTableRelate{}).Create(map[string]interface{}{
		"database_id": databaseId,
		"table_id":    tableId,
		"state":       1,
		"create_time": util.GetUnixTime(),
	}).Error
	return err
}
