package model

import (
	"commons/util"
	"store-service/dao"
)

type TableUserRelate struct {
	TableUserId int   `json:"table_user_id" gorm:"column:table_user_id;primary_key;AUTO_INCREMENT"`
	TableId     int   `json:"table_id" gorm:"column:table_id"`
	UserId      int   `json:"user_id" gorm:"column:user_id"`
	State       int   `json:"state" gorm:"column:state"`
	CreateTime  int64 `json:"create_time" gorm:"column:create_time"`
	UpdateTime  int64 `json:"update_time" gorm:"column:update_time"`
}

func (TableUserRelate) TableName() string {
	return "t_table_user_relate"
}

func InitTableUser(tableId int, userId int) error {
	var t = TableUserRelate{
		TableId:    tableId,
		UserId:     userId,
		State:      1,
		CreateTime: util.GetUnixTime(),
		UpdateTime: 0,
	}
	err := dao.System.Select("table_id", "user_id", "state", "create_time").Create(&t).Error
	return err
}
