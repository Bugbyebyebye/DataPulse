package model

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
