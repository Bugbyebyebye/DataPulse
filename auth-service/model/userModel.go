package model

import (
	"auth-service/dao"
	"commons/util"
)

//用户登录注册相关数据库操作

type User struct {
	UserId     int    `gorm:"column:user_id;primary_key;AUTO_INCREMENT"`
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	Email      string `gorm:"column:email"`
	Role       string `gorm:"column:role"`
	State      int    `gorm:"column:state"`
	CreateTime int64  `gorm:"column:create_time"`
	UpdateTime int64  `gorm:"column:update_time"`
}

func (User) TableName() string {
	return "t_user"
}

// GetUserByUsername 查询用户名是否重复
func GetUserByUsername(username string) (User, error) {
	var user User
	err := dao.Db.Where("username = ?", username).First(&user).Error
	return user, err
}

// GetUserByEmail 根据邮箱查询用户
func GetUserByEmail(email string) (User, error) {
	var user User
	err := dao.Db.Where("email = ?", email).First(&user).Error
	return user, err
}

// InsertUser 新增用户数据
func InsertUser(username string, password string, email string) (int, error) {
	user := User{Username: username, Password: password, Email: email, Role: "user", State: 0, CreateTime: util.GetUnixTime()}
	create := dao.Db.Select("username", "password", "email", "create_time").Create(&user)
	return user.UserId, create.Error
}
