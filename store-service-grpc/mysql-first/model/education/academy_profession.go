package model

type AcademyProfessionRelate struct {
	AcademyProfessionId int    `json:"academy_profession_id" gorm:"column:academy_profession_id;primary_key;AUTO_INCREMENT"`
	AcademyId           int    `json:"academy_id" gorm:"column:academy_id"`
	ProfessionId        string `json:"profession_id" gorm:"column:profession_id"`
	State               int    `json:"-" gorm:"column:state"`
	CreateTime          int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime          int64  `json:"update_time" gorm:"column:update_time"`
}

func (AcademyProfessionRelate) TableName() string {
	return "t_academy_profession_relate"
}
