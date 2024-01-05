package handle

import (
	"auth-service/model"
	"commons/logsmodel"
	"commons/result"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var res result.Result

type UserListRes struct {
	UserId     int    `json:"user_id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	Authority  int    `json:"authority"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

// GetUserList 获取用户列表
func (*AuthHandler) GetUserList(ctx *gin.Context) {

	id := ctx.GetInt("id")
	role := ctx.GetString("role")
	authority := ctx.GetInt("authority")

	if id == 0 && role != "admin" && authority == 0 {
		ctx.JSON(http.StatusOK, res.Fail(4001, "抱歉，权限不够！"))
		return
	}

	var userList []UserListRes
	list, err := model.GetUserList()
	for _, user := range list {
		var userRes UserListRes
		userRes.UserId = user.UserId
		userRes.Email = user.Email
		userRes.Username = user.Username
		userRes.Authority = user.Authority
		userRes.Role = user.Role
		userRes.CreateTime = user.CreateTime
		userRes.UpdateTime = user.UpdateTime
		userList = append(userList, userRes)
	}
	if err != nil {
		log.Printf("err => %s", err)
		ctx.JSON(http.StatusOK, res.Fail(400, "查询失败"))
		return
	}

	ctx.JSON(http.StatusOK, res.Success(userList))
}

// SetUserRole 设置用户角色
func (*AuthHandler) SetUserRole(ctx *gin.Context) {
	//获取用户id
	idStr := ctx.PostForm("id")
	role := ctx.PostForm("role")
	userId, _ := strconv.Atoi(idStr)
	err := model.SetUserRole(userId, role)
	if err != nil {
		ctx.JSON(http.StatusOK, res.Fail(4001, "修改错误"))
		logsmodel.PostActionLogs(userId, "修改用户角色", "Failed")
		return
	}

	ctx.JSON(http.StatusOK, res.Success("角色设置成功"))
	logsmodel.PostActionLogs(userId, "修改用户角色", "Success")
}

// SetUserAuthority 设置用户权限
func (*AuthHandler) SetUserAuthority(ctx *gin.Context) {
	idStr := ctx.PostForm("id")
	authStr := ctx.PostForm("authority")
	id, _ := strconv.Atoi(idStr)
	auth, _ := strconv.Atoi(authStr)

	err := model.SetUserAuthority(id, auth)
	if err != nil {
		logsmodel.PostActionLogs(id, "修改用户权限", "Failed")
		ctx.JSON(http.StatusOK, res.Fail(4001, "修改错误"))
		return
	}
	logsmodel.PostActionLogs(id, "修改用户权限", "Failed")
	ctx.JSON(http.StatusOK, res.Success("权限设置成功"))
}

// DeleteUser 删除用户
func (*AuthHandler) DeleteUser(ctx *gin.Context) {
	idStr := ctx.PostForm("id")
	stateStr := ctx.PostForm("state")
	id, _ := strconv.Atoi(idStr)
	state, _ := strconv.Atoi(stateStr)

	err := model.DeleteUser(id, state)
	if err != nil {
		logsmodel.PostActionLogs(id, "删除用户", "Failed")
		ctx.JSON(http.StatusOK, res.Fail(4001, "删除错误"))
		return
	}
	logsmodel.PostActionLogs(id, "删除用户", "Success")

	ctx.JSON(http.StatusOK, res.Success("删除成功"))
}
