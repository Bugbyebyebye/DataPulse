package model

type BookStudentRelate struct {
	BorrowId        int    `json:"boorow_id" gorm:"column:borrow_id;primary_key;AUTO_INCREMENT"`
	StudentId       int    `json:"student_id" gorm:"column:student_id"`
	StudentName     string `json:"student_name" gorm:"column:student_name"`
	BookId          int    `json:"book_id" gorm:"column:book_id"`
	BorrowStartTime int64  `json:"borrow_start_time" gorm:"column:borrow_start_time"`
	BorrowEndTime   int64  `json:"borrow_end_time" gorm:"column:borrow_end_time"`
	State           int    `json:"-" gorm:"column:state"`
	CreateTime      int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime      int64  `json:"update_time" gorm:"column:update_time"`
}

func (BookStudentRelate) TableName() string {
	return "t_borrow_relate"
}
