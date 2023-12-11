package model

type ClassStudentRelate struct {
	ClassStudentId int    `json:"class_student_id" gorm:"column:class_student_id;primary_key;AUTO_INCREMENT"`
	ClassId        int    `json:"class_id" gorm:"column:class_id"`
	StudentId      string `json:"student_id" gorm:"column:student_id"`
	State          int    `json:"-" gorm:"column:state"`
	CreateTime     int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime     int64  `json:"update_time" gorm:"column:update_time"`
}

func (ClassStudentRelate) TableName() string {
	return "t_class_student_relate"
}
