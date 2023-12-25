package model

import (
	"commons/util"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log-service/dao"
)

// SearchUserLogs 通过用户查询日志
func SearchUserLogs(userid int, table string) ([]map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = ? AND state != 0", table)
	// 执行查询
	rows, err := dao.Db.Query(query, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 获取列名
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	logs := make([]map[string]interface{}, 0)

	// 迭代查询结果
	for rows.Next() {
		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))
		for i := range values {
			scanArgs[i] = &values[i]
		}

		// 扫描行数据
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		entry := make(map[string]interface{})
		// 将行数据存储在映射中
		for i, col := range values {
			if col != nil {
				entry[columns[i]] = col
			} else {
				entry[columns[i]] = nil
			}
		}
		logs = append(logs, entry)
	}
	//解码切片为utf-8格式
	if err = rows.Err(); err != nil {
		return nil, err
	}
	for _, logEntry := range logs {
		for key, value := range logEntry {
			// 检查值是否是字节切片类型
			if byteSlice, ok := value.([]byte); ok {
				// 将字节切片转换为字符串
				strValue := string(byteSlice)
				logEntry[key] = strValue
			}
		}
	}
	//fmt.Println(logs)
	return logs, nil
}

/*
数据库删除操作
*/

// DeleteUserLogs 删除用户相关的所有日志
func DeleteUserLogs(userid int, table string) error {
	// 确保连接正常
	err := dao.Db.Ping()
	if err != nil {
		fmt.Printf("数据库不健康 : %s\n", err.Error())
		return err
	}
	query := fmt.Sprintf("UPDATE %s SET state = 0 WHERE user_id = ?", table)
	_, err = dao.Db.Exec(query, userid)
	return err
}

// DeleteOneLogs 通过id删除日志
func DeleteOneLogs(id int, table string) error {
	// 确保连接正常
	err := dao.Db.Ping()
	if err != nil {
		fmt.Printf("数据库不健康 : %s\n", err.Error())
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET state = 0 WHERE id = ?", table)
	stmt, err := dao.Db.Prepare(query)
	if err != nil {
		fmt.Printf("无法准备查询 : %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Printf("执行查询时出错 : %s\n", err.Error())
		return err
	}

	return nil
}

/*
数据库插入操作
*/

// RecordUserActionLog 插入用户操作日志记录数据库
func RecordUserActionLog(UserId, Content, Status string) error {
	query := `INSERT INTO action_logs_info (user_id, content, status, update_time, state) VALUES (?, ?, ?, ?, ?)`
	// 执行插入操作
	logDate := util.GetUnixTime()
	State := 1
	_, err := dao.Db.Exec(query, UserId, Content, Status, logDate, State)
	return err
}

// RecordUserTackLog 插入任务调度日志记录数据库
func RecordUserTackLog(UserId, TaskName, TaskStatus, Takingtime string) error {
	query := `INSERT INTO task_log_info (user_id, task_name, task_status ,update_time, state, taking_time) VALUES (?, ?, ?, ?, ?, ?)`
	// 执行插入操作
	logDate := util.GetUnixTime()
	State := 1
	_, err := dao.Db.Exec(query, UserId, TaskName, TaskStatus, logDate, State, Takingtime)
	return err
}

// RecordUserInvokeLog 插入api调用日志记录数据库
func RecordUserInvokeLog(ApiUrl, UserId, Invokeip, Invokelog, Status string) error {
	query := `INSERT INTO invoke_log_info (api_url, state, user_id, status, invoke_time, invoke_ip, invoke_log) VALUES (?, ?, ?, ?, ?, ?, ?)`
	// 执行插入操作
	logDate := util.GetUnixTime()
	State := 1
	_, err := dao.Db.Exec(query, ApiUrl, State, UserId, Status, logDate, Invokeip, Invokelog)
	return err
}
