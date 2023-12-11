package model

type RoomStudentRelate struct {
	AppointId         int    `json:"appoint_id" gorm:"column:appoint_id;primary_key;AUTO_INCREMENT"`
	AppointPersonId   int    `json:"appoint_person_id" gorm:"column:appoint_person_id"`
	AppointPersonName string `json:"appoint_person_name" gorm:"column:appoint_person_name"`
	RoomId            int    `json:"room_id" gorm:"column:room_id"`
	AppointStartTime  int64  `json:"appoint_start_time" gorm:"column:appoint_start_time"`
	AppointEndTime    int64  `json:"appoint_end_time" gorm:"column:appoint_end_time"`
	State             int    `json:"-" gorm:"column:state"`
	CreateTime        int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime        int64  `json:"update_time" gorm:"column:update_time"`
}

func (RoomStudentRelate) TableName() string {
	return "t_appoint_relate"
}
