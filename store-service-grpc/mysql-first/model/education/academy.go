package model

type AcademyInfo struct {
	AcademyId     int    `json:"academy_id" gorm:"column:academy_id;primary_key;AUTO_INCREMENT"`
	AcademyName   int    `json:"academy_name" gorm:"column:academy_name"`
	AcademyDesc   string `json:"academy_desc" gorm:"column:academy_desc"`
	ProfessionNum string `json:"profession_num" gorm:"column:profession_num"`
	State         int    `json:"-" gorm:"column:state"`
	CreateTime    int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime    int64  `json:"update_time" gorm:"column:update_time"`
}

func (AcademyInfo) TableName() string {
	return "t_academy_info"
}
