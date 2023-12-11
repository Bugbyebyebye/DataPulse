package model

type TeacherInfo struct {
	TeacherId     int    `json:"teacher_id" gorm:"column:teacher_id;primary_key;AUTO_INCREMENT"`
	TeacherName   int    `json:"teacher_name" gorm:"column:user_id"`
	TeacherAge    string `json:"teacher_age" gorm:"column:nickname"`
	TeacherGender string `json:"teacher_gender" gorm:"column:desc"`
	TeacherApart  string `json:"teacher_apart" gorm:"column:avatar"`
	State         int    `json:"-" gorm:"column:state"`
	CreateTime    int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime    int64  `json:"update_time" gorm:"column:update_time"`
}

func (TeacherInfo) TableName() string {
	return "t_teacher_info"
}
