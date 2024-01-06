package model

import (
	"commons/util"
	"log"
	"store-service/config"
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
	err := config.System.Select("table_id", "user_id", "state", "create_time").Create(&t).Error
	return err
}

//func Update(tableName string) error {
//	var dur DatabaseUserRelate
//	dao.Warehouse.Model(&dur).Where("table_name = ?", tableName).Updates(map[string]interface{}{
//		"UpdateTime": util.GetUnixTime(),
//	})
//}

// GetTableIdList 根据用户id 获取用户数据表id列表
func GetTableIdList(userId int) ([]int, error) {
	log.Printf("userId => %v", userId)
	var result []int
	var tur TableUserRelate
	err := config.System.Model(&tur).Where("user_id = ? and state = 1", userId).Select("table_id").Find(&result).Error
	return result, err
}

// DeleteUserTableRelate 逻辑删除用户数据表关联
func DeleteUserTableRelate(userId int, tableId int) error {
	err := config.System.Model(TableUserRelate{}).Where("user_id = ? and table_id = ?", userId, tableId).Updates(map[string]interface{}{
		"state":       0,
		"update_time": util.GetUnixTime(),
	}).Error

	return err
}
