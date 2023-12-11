package model

type DatabaseUserRelate struct {
	DatabaseUserId int   `json:"database_user_id" gorm:"column:database_user_id;primary_key;AUTO_INCREMENT"`
	DatabaseId     int   `json:"database_id" gorm:"column:database_id"`
	UserId         int   `json:"user_id" gorm:"column:user_id"`
	State          int   `json:"state" gorm:"column:state"`
	CreateTime     int64 `json:"create_time" gorm:"column:create_time"`
	UpdateTime     int64 `json:"update_time" gorm:"column:update_time"`
}

func (DatabaseUserRelate) TableName() string {
	return "t_database_user_relate"
}
