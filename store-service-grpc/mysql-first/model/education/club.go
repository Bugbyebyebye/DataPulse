package model

type ClubInfo struct {
	ClubId       int    `json:"club_id" gorm:"column:club_id;primary_key;AUTO_INCREMENT"`
	ClubName     int    `json:"club_name" gorm:"column:club_name"`
	ClubHeaderId string `json:"club_header_id" gorm:"column:club_header_id"`
	PersonNum    string `json:"person_num" gorm:"column:person_num"`
	State        int    `json:"-" gorm:"column:state"`
	CreateTime   int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime   int64  `json:"update_time" gorm:"column:update_time"`
}

func (ClubInfo) TableName() string {
	return "t_club_info"
}
