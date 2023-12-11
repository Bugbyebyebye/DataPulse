package model

type RoomInfo struct {
	RoomId       int    `json:"room_id" gorm:"column:room_id;primary_key;AUTO_INCREMENT"`
	RoomLocation int    `json:"room_location" gorm:"column:room_location"`
	RoomDesc     string `json:"room_desc" gorm:"column:room_desc"`
	State        int    `json:"-" gorm:"column:state"`
	CreateTime   int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime   int64  `json:"update_time" gorm:"column:update_time"`
}

func (RoomInfo) TableName() string {
	return "t_room_info"
}
