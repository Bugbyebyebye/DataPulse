package model

type ProfessionInfo struct {
	ProfessionId   int    `json:"profession_id" gorm:"column:profession_id;primary_key;AUTO_INCREMENT"`
	ProfessionName int    `json:"profession_name" gorm:"column:profession_name"`
	ProfessionDesc string `json:"profession_desc" gorm:"column:profession_desc"`
	ClassNum       string `json:"class_num" gorm:"column:class_name"`
	State          int    `json:"-" gorm:"column:state"`
	CreateTime     int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime     int64  `json:"update_time" gorm:"column:update_time"`
}

func (ProfessionInfo) TableName() string {
	return "t_profession_info"
}
