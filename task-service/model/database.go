package model

import (
	"commons/util"
	"fmt"
	"task-service/dao"
)

// APIInfo API数据库结构体
type APIInfo struct {
	ID          int    `json:"api_id"`
	Name        string `json:"api_name"`
	URL         string `json:"api_url"`
	Description string `json:"api_desc"`
	State       int    `json:"state"`
	CreateTime  int64  `json:"create_time"`
	UpdateTime  int64  `json:"update_time"`
}

// 将api插入到数据库中

func InsetAPIList(APIName, APIUrl, APIDesc string, UserID int) error {
	state := 1
	CreatTime := util.GetUnixTime()

	query := fmt.Sprintf(
		"INSERT INTO `%s` SET `state` = %d, `create_time` = %d, `api_name` = '%s', `api_url` = '%s', `api_desc` = '%s'",
		"t_api_info",
		state,
		CreatTime,
		APIName,
		APIUrl,
		APIDesc,
	)
	result, err := dao.Db.Exec(query)
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	query = fmt.Sprintf(
		"INSERT INTO `%s` SET `state` = %d, `create_time` = %d, `api_id` = %d, `user_id` = %d",
		"t_api_user_relate",
		state,
		CreatTime,
		lastID,
		UserID,
	)
	_, err = dao.Db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// 用id删除API/动态删除

func DeleteAPI(UserID, APID int) error {
	if APID == 0 {
		// 如果 APID 为 0，删除所有具有特定 UserID 的记录
		query := fmt.Sprintf("DELETE FROM t_api_user_relate WHERE user_id = %d", UserID)
		_, err := dao.Db.Exec(query)
		query = fmt.Sprintf("DELETE FROM t_api_info WHERE user_id = %d", UserID)
		_, err = dao.Db.Exec(query)
		return err
	}

	// 删除具有特定 UserID 和 APID 的记录
	query := fmt.Sprintf("DELETE FROM t_api_user_relate WHERE user_id = %d AND api_id = %d", UserID, APID)
	_, err := dao.Db.Exec(query)
	// 删除具有特定 UserID 和 APID 的记录
	query = fmt.Sprintf("DELETE FROM t_api_info WHERE user_id = %d AND api_id = %d", UserID, APID)
	_, err = dao.Db.Exec(query)
	return err
}

// SearchAPIList 通过用户id查询接口id列表
func SearchAPIList(UserID int) ([]APIInfo, error) {
	// 查询 t_api_user_relate 获取所有的 api_id
	query := fmt.Sprintf("SELECT api_id FROM t_api_user_relate WHERE user_id = %d", UserID)
	rows, err := dao.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 保存查询到的所有 api_id
	var apiIDs []int
	for rows.Next() {
		var apiID int
		if err := rows.Scan(&apiID); err != nil {
			return nil, err
		}
		apiIDs = append(apiIDs, apiID)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// 查询 t_api_info 获取对应的 API 信息
	var apiList []APIInfo
	for _, apiID := range apiIDs {
		query := fmt.Sprintf("SELECT * FROM t_api_info WHERE api_id = %d", apiID)
		rows, err := dao.Db.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		// 处理查询结果并添加到 apiList
		for rows.Next() {
			var api APIInfo
			if err := rows.Scan(&api.ID, &api.Name, &api.URL, &api.Description, &api.State, &api.CreateTime, &api.UpdateTime); err != nil {
				return nil, err
			}
			apiList = append(apiList, api)
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}
	}
	return apiList, nil
}

// TODO 通过接口ID查询接口信息
