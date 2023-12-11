package model

type BookInfo struct {
	BookId      int    `json:"book_id" gorm:"column:book_id;primary_key;AUTO_INCREMENT"`
	BookName    int    `json:"book_name" gorm:"column:book_name"`
	PublishTime string `json:"publish_time" gorm:"column:publish_time"`
	BookDesc    string `json:"book_desc" gorm:"column:book_desc"`
	BookAuthor  string `json:"book_author" gorm:"column:book_author"`
	State       int    `json:"-" gorm:"column:state"`
	CreateTime  int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime  int64  `json:"update_time" gorm:"column:update_time"`
}

func (BookInfo) TableName() string {
	return "t_book_info"
}
