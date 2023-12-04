package model

import (
	"auth-service/dao"
	"commons/util"
)

//用户信息相关操作

type Info struct {
	UserInfoId int    `gorm:"column:user_info_id;primary_key;AUTO_INCREMENT"`
	UserId     int    `gorm:"column:user_id"`
	Nickname   string `gorm:"column:nickname"`
	Desc       string `gorm:"column:desc"`
	Avatar     string `gorm:"column:avatar"`
	CreateTime int64  `gorm:"column:create_time"`
	UpdateTime int64  `gorm:"column:update_time"`
}

func (Info) TableName() string {
	return "t_user_info"
}

// GetUserInfoByUserId 通过用户id获取用户信息
func GetUserInfoByUserId(userId int) (Info, error) {
	var info Info
	err := dao.Db.Where("user_id = ?", userId).First(&info).Error
	return info, err
}

// InitUserInfo 初始化用户信息
func InitUserInfo(userId int) error {
	info := Info{UserId: userId, Nickname: "默认昵称", Desc: "暂无个人简介", CreateTime: util.GetUnixTime()}
	create := dao.Db.Select("user_id", "nickname", "desc", "create_time").Create(&info)
	return create.Error
}

// SetUserInfo 更新用户信息
func SetUserInfo(userId int, nickname string, desc string, avatar string) error {
	var info Info
	err := dao.Db.Model(&info).Where("user_id = ?", userId).Updates(map[string]interface{}{
		"Nickname":   nickname,
		"Desc":       desc,
		"Avatar":     avatar,
		"UpdateTime": util.GetUnixTime(),
	}).Error
	return err
}
