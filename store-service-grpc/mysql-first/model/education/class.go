package model

type ClassInfo struct {
	ClassId    int    `json:"class_id" gorm:"column:class_id;primary_key;AUTO_INCREMENT"`
	ClassName  int    `json:"class_name" gorm:"column:class_name"`
	ClassDesc  string `json:"class_desc" gorm:"column:class_desc"`
	StudentNum string `json:"student_num" gorm:"column:student_num"`
	State      int    `json:"-" gorm:"column:state"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time"`
}

func (ClassInfo) TableName() string {
	return "t_class_info"
}
