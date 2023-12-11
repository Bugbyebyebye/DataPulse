package model

type ProfessionClassRelate struct {
	ProfessionClassId int    `json:"profession_class_id" gorm:"column:profession_class_id;primary_key;AUTO_INCREMENT"`
	ClassId           int    `json:"class_id" gorm:"column:class_id"`
	ProfessionId      string `json:"profession_id" gorm:"column:profession_id"`
	State             int    `json:"-" gorm:"column:state"`
	CreateTime        int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime        int64  `json:"update_time" gorm:"column:update_time"`
}

func (ProfessionClassRelate) TableName() string {
	return "t_profession_class_relate"
}
