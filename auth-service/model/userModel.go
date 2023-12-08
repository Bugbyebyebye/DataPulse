package model

import (
	"auth-service/dao"
	"commons/util"
)

//用户登录注册相关数据库操作

type User struct {
	UserId     int    `gorm:"column:user_id;primary_key;AUTO_INCREMENT"` //用户id
	Username   string `gorm:"column:username"`                           //用户名
	Password   string `gorm:"column:password"`                           //密码
	Email      string `gorm:"column:email"`
	Role       string `gorm:"column:role"`
	Authority  int    `gorm:"column:authority"`
	State      int    `gorm:"column:state"`
	CreateTime int64  `gorm:"column:create_time"`
	UpdateTime int64  `gorm:"column:update_time"`
}

func (User) TableName() string {
	return "t_user"
}

// GetUserById 根据用户id查询用户信息
func GetUserById(userId int) (User, error) {
	var user User
	err := dao.Db.Where("user_id = ?", userId).First(&user).Error
	return user, err
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

// GetUserList 获取用户列表
func GetUserList() ([]User, error) {
	var userList []User
	err := dao.Db.Where("state = ?", 1).Find(&userList).Error
	if err != nil {
		return nil, err
	}
	return userList, nil
}

// GetUserByAccount 通过用户id获取用户信息
func GetUserByAccount(userId int) (User, error) {
	var user User
	err := dao.Db.Select("user_id, username, email, create_time, update_time").Where("user_id = ?", userId).First(&user).Error
	return user, err
}

// InitUser 新增用户数据
func InitUser(username string, password string, email string) (int, error) {
	user := User{Username: username, Password: password, Email: email, Role: "user", Authority: 0, State: 0, CreateTime: util.GetUnixTime()}
	create := dao.Db.Create(&user)
	return user.UserId, create.Error
}

// SetAccount 更新用户账号信息(邮箱，账号，密码)
func SetAccount(userId int, Username string, Password string, Email string) error {
	var user User
	err := dao.Db.Model(&user).Where("user_id = ?", userId).Updates(map[string]interface{}{
		"Username":   Username,
		"Password":   Password,
		"Email":      Email,
		"UpdateTime": util.GetUnixTime(),
	}).Error
	return err
}

// SetUserRole 设置用户角色
func SetUserRole(userId int, role string) error {
	var user User
	err := dao.Db.Model(&user).Where("user_id = ?", userId).Updates(map[string]interface{}{
		"Role":       role,
		"UpdateTime": util.GetUnixTime(),
	}).Error
	return err
}

// SetUserAuthority 设置用户权限
func SetUserAuthority(userId int, authority int) error {
	var user User
	err := dao.Db.Model(&user).Where("user_id = ?", userId).Updates(map[string]interface{}{
		"Authority":  authority,
		"UpdateTime": util.GetUnixTime(),
	}).Error
	return err
}

// DeleteUser 逻辑删除用户记录
func DeleteUser(userId int) error {
	var user User
	err := dao.Db.Model(&user).Where("user_id = ?", userId).Updates(map[string]interface{}{
		"State":      0,
		"UpdateTime": util.GetUnixTime(),
	}).Error
	return err
}
