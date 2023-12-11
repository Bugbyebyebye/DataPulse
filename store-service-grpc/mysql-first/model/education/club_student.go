package model

type ClubStudentRelate struct {
	ClubStudentId int    `json:"club_student_id" gorm:"column:club_student_id;primary_key;AUTO_INCREMENT"`
	ClubId        int    `json:"club_id" gorm:"column:club_id"`
	StudentId     string `json:"student_id" gorm:"column:student_id"`
	State         int    `json:"-" gorm:"column:state"`
	CreateTime    int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime    int64  `json:"update_time" gorm:"column:update_time"`
}

func (ClubStudentRelate) TableName() string {
	return "t_club_student_relate"
}
