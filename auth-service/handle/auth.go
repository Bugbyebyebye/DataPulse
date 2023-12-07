package handle

import (
	"auth-service/dao"
	"auth-service/model"
	"auth-service/util"
	"commons/result"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
)

type LoginReq struct {
	Type     int    `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     int    `json:"code"`
}

type LoginRes struct {
	Token     string `json:"token"`
	UserId    int    `json:"user_id"`
	Role      string `json:"role"`
	Authority int    `json:"authority"`
}

// UserLogin 用户登录
func (*AuthHandler) UserLogin(ctx *gin.Context) {
	res := &result.Result{}
	req := &LoginReq{}

	err := ctx.BindJSON(req)
	if err != nil {
		log.Printf("json数据错误 err => %s", err)
		ctx.JSON(200, res.Fail(400, err.Error()))
		return
	}

	log.Printf("req => %+v\n", req)

	if req.Type == 1 {
		//邮箱验证码登录
		codeStr, err := dao.Rc.Get(context.Background(), "DATAPULSE"+req.Email)
		if err != nil {
			log.Printf("redis error => %s", err)
			if errors.Is(err, redis.ErrNil) {
				ctx.JSON(200, res.Fail(4001, "验证码已过期，请重新发送"))
			}
			return
		}

		//校验验证码
		code, _ := strconv.Atoi(codeStr)
		if req.Code == code {
			user, err := model.GetUserByEmail(req.Email)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					ctx.JSON(200, res.Fail(4001, "登录失败，用户不存在！"))
					return
				}
			}
			token, err := util.CreateToken(user.UserId, user.Username, user.Role, user.Authority)
			if err != nil {
				log.Printf("token err => %s", err)
			}

			var data LoginRes
			data.UserId = user.UserId
			data.Token = token
			data.Role = user.Role
			data.Authority = user.Authority

			ctx.JSON(http.StatusOK, res.Success(data))
			return
		} else {
			ctx.JSON(200, res.Fail(4001, "验证码错误！"))
			return
		}
	} else if req.Type == 2 {
		//账号密码登录
		user, err := model.GetUserByUsername(req.Username)
		if err != nil {
			log.Printf("getUserByUsername error => %s", err)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(200, res.Fail(4001, "登录失败，用户不存在！"))
				return
			}
		}

		if user.Password == req.Password {
			//登录成功
			token, err := util.CreateToken(user.UserId, user.Username, user.Role, user.Authority)
			if err != nil {
				log.Printf("err => %s", err)
			}

			var data LoginRes
			data.UserId = user.UserId
			data.Token = token
			data.Role = user.Role
			data.Authority = user.Authority

			ctx.JSON(http.StatusOK, res.Success(data))

		} else {
			//校验密码错误
			ctx.JSON(200, res.Fail(4001, "密码错误，登录失败！"))
		}
	}
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     int    `json:"code"`
}

// UserRegister 用户注册
func (*AuthHandler) UserRegister(ctx *gin.Context) {
	res := &result.Result{}
	req := &RegisterReq{}

	err := ctx.BindJSON(req)
	if err != nil {
		log.Printf("json数据错误 err => %s", err)
		ctx.JSON(http.StatusOK, res.Fail(400, "json数据错误！"))
		return
	}

	log.Printf("req => %+v\n", req)
	//判断参数是否为空
	if req.Code == 0 && req.Email == "" && req.Username == "" && req.Password == "" {
		ctx.JSON(http.StatusOK, res.Fail(400, "参数不能为空！"))
		return
	}

	//从redis获取验证码
	codeStr, err := dao.Rc.Get(context.Background(), "DATAPULSE"+req.Email)
	if err != nil {
		log.Printf("redis error => %s", err)
		ctx.JSON(http.StatusOK, res.Fail(400, "验证码错误！"))
		return
	}
	code, _ := strconv.Atoi(codeStr)

	//校验验证码
	if req.Code == code {
		//邮箱验证码校验成功
		_, err := model.GetUserByUsername(req.Username)
		if err == nil {
			ctx.JSON(200, res.Fail(4001, "该用户名已注册，请重新输入！"))
			return
		}

		id, err := model.InitUser(req.Username, req.Password, req.Email)
		if err != nil {
			log.Printf("mysql error => %s", err)
			return
		}
		err = model.InitUserInfo(id)

		ctx.JSON(200, res.Success("用户注册成功！"))
		return
	} else {
		ctx.JSON(200, res.Fail(4001, "验证码输入错误！"))
		return
	}
}

// UserInfoRes 用户信息
type UserInfoRes struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Nickname   string `json:"nickname"`
	Desc       string `json:"desc"`
	CreateTime int64  `json:"create_time"`
	Avatar     string `json:"avatar"`
}

// GetUserInfo 获取个人中心信息
func (*AuthHandler) GetUserInfo(ctx *gin.Context) {
	res := &result.Result{}

	token := ctx.Request.Header.Get("token")
	claims, _ := util.ParseToken(token)
	id := claims.Id
	log.Printf("id => %v", id)
	//username, ok := ctx.Get("username")
	if id == 0 {
		ctx.JSON(200, res.Fail(4001, "用户id获取失败"))
		return
	}

	info, err := model.GetUserInfoByUserId(id)
	if err != nil {
		log.Printf("mysql error => %s", err)
		return
	}
	user, err := model.GetUserById(id)
	if err != nil {
		log.Printf("mysql error => %s", err)
		return
	}

	var data UserInfoRes
	data.Username = user.Username
	data.Email = user.Email
	data.Nickname = info.Nickname
	data.Desc = info.Desc
	data.Avatar = info.Avatar
	data.CreateTime = user.CreateTime

	ctx.JSON(http.StatusOK, res.Success(data))
}

type InfoReq struct {
	Nickname string `json:"nickname"`
	Desc     string `json:"desc"`
	Avatar   string `json:"avatar"`
}

// SetUserInfo 设置个人中心信息
func (*AuthHandler) SetUserInfo(ctx *gin.Context) {
	res := &result.Result{}
	req := &InfoReq{}
	log.Printf("req => %s", req)

	token := ctx.Request.Header.Get("token")
	claims, _ := util.ParseToken(token)
	id := claims.Id

	err := ctx.BindJSON(req)
	if err != nil {
		log.Printf("json error => %s", err)
		return
	}
	log.Printf("req => %+v\n", req)

	err = model.SetUserInfo(id, req.Nickname, req.Desc, req.Avatar)
	if err != nil {
		log.Printf("mysql error => %s", err)
		ctx.JSON(200, res.Fail(4001, "更新失败！"))
		return
	}
	user, err := model.GetUserInfoByUserId(id)
	ctx.JSON(http.StatusOK, res.Success(user))
}

// SetAccount 更改账号信息
func (h *AuthHandler) SetAccount(ctx *gin.Context) {
	res := &result.Result{}
	token := ctx.Request.Header.Get("token")
	claims, _ := util.ParseToken(token)
	id := claims.Id
	req := &RegisterReq{}

	err := ctx.BindJSON(req)
	if err != nil {
		log.Printf("json数据错误 err => %s", err)
		ctx.JSON(http.StatusOK, res.Fail(400, "json数据错误！"))
		return
	}
	log.Printf("req %s", req)
	//从redis获取验证码
	codeStr, err := dao.Rc.Get(context.Background(), "DATAPULSE"+req.Email)
	if err != nil {
		log.Printf("redis error => %s", err)
	}
	code, _ := strconv.Atoi(codeStr)

	//校验验证码
	if req.Code == code {
		//邮箱验证码校验成功
		//如果通过token获取的username不等于username的话，重复之后return
		user, err := model.GetUserByAccount(id)
		if !(user.Username == req.Username) {
			_, err := model.GetUserByUsername(req.Username)
			if err == nil {
				ctx.JSON(200, res.Fail(4001, "用户名重复,请更改用户名！"))
				return
			}
		}

		err = model.SetAccount(id, req.Username, req.Password, req.Email)
		if err != nil {
			log.Printf("mysql error => %s", err)
			return
		}
		ctx.JSON(200, res.Success("更新成功！"))
		return
	} else {
		ctx.JSON(200, res.Fail(4001, "验证码输入错误！"))
		return
	}
}
