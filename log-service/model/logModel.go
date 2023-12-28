package model

import (
	"commons/util"
	_ "github.com/go-sql-driver/mysql"
	"log-service/dao"
)

/*
数据库插入操作
*/

// RecordUserActionLog 插入用户操作日志记录数据库
func RecordUserActionLog(UserId int, Content, Status string) error {
	query := `INSERT INTO action_logs_info (user_id, content, status, update_time, state) VALUES (?, ?, ?, ?, ?)`
	// 执行插入操作
	logDate := util.GetUnixTime()
	State := 1
	_, err := dao.Db.Exec(query, UserId, Content, Status, logDate, State)
	return err
}

// RecordUserTackLog 插入任务调度日志记录数据库
func RecordUserTackLog(UserId int, TaskName, TaskStatus, Takingtime string) error {
	query := `INSERT INTO task_log_info (user_id, task_name, task_status ,update_time, state, taking_time) VALUES (?, ?, ?, ?, ?, ?)`
	// 执行插入操作
	logDate := util.GetUnixTime()
	State := 1
	_, err := dao.Db.Exec(query, UserId, TaskName, TaskStatus, logDate, State, Takingtime)
	return err
}

// RecordUserInvokeLog 插入api调用日志记录数据库
func RecordUserInvokeLog(ApiUrl string, UserId int, Invokeip, Invokelog, Status string) error {
	query := `INSERT INTO invoke_log_info (api_url, state, user_id, status, invoke_time, invoke_ip, invoke_log) VALUES (?, ?, ?, ?, ?, ?, ?)`
	// 执行插入操作
	logDate := util.GetUnixTime()
	State := 1
	_, err := dao.Db.Exec(query, ApiUrl, State, UserId, Status, logDate, Invokeip, Invokelog)
	return err
}
