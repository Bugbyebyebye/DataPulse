package model

/**
t_student_info
学生信息表
*/

type StudentInfo struct {
	StudentId      int    `json:"student_id" gorm:"column:student_id;primary_key;AUTO_INCREMENT"`
	StudentName    string `json:"student_name" gorm:"column:student_name"`
	StudentGender  string `json:"student_gender" gorm:"column:student_gender"`
	StudentAge     string `json:"student_age" gorm:"column:student_age"`
	StudentPhone   string `json:"student_phone" gorm:"column:student_phone"`
	StudentClass   string `json:"student_class" gorm:"column:student_class"`
	StudentAcademy string `json:"student_academy" gorm:"column:student_academy"`
	State          int    `json:"-" gorm:"column:state"`
	CreateTime     int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime     int64  `json:"update_time" gorm:"column:update_time"`
}

func (StudentInfo) TableName() string {
	return "t_student_info"
}
