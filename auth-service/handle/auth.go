package handle

import (
	"auth-service/config"
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

type AuthHandler struct {
	cache config.Cache
}

func New() *AuthHandler {
	return &AuthHandler{
		cache: dao.Rc,
	}
}

type LoginReq struct {
	Type     int    `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     int    `json:"code"`
}

// UserLogin 用户登录
func (h *AuthHandler) UserLogin(ctx *gin.Context) {
	res := &result.Result{}
	req := &LoginReq{}

	err := ctx.BindJSON(req)
	if err != nil {
		log.Printf("json数据错误 err => %s", err)
		ctx.JSON(200, res.Fail(400, err.Error()))
		return
	}

	log.Printf("req => %s", req)

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
			token, err := util.CreateToken(user.UserId, user.Username)
			if err != nil {
				log.Printf("token err => %s", err)
			}
			ctx.JSON(200, res.Success(token))
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
			token, err := util.CreateToken(user.UserId, user.Username)
			if err != nil {
				log.Printf("err => %s", err)
			}
			ctx.JSON(http.StatusOK, res.Success(token))

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
func (h *AuthHandler) UserRegister(ctx *gin.Context) {
	res := &result.Result{}
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
		id, err := model.InsertUser(req.Username, req.Password, req.Email)
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

// GetUserInfo 获取个人中心信息
func (h *AuthHandler) GetUserInfo(ctx *gin.Context) {
	res := &result.Result{}

	token := ctx.Request.Header.Get("token")
	claims, _ := util.ParseToken(token)
	id := claims.Id

	//username, ok := ctx.Get("username")
	if id == 0 {
		ctx.JSON(200, res.Fail(4001, "用户id获取失败"))
		return
	}

	user, err := model.GetUserInfoByUserId(id)
	if err != nil {
		log.Printf("mysql error => %s", err)
		return
	}

	ctx.JSON(http.StatusOK, res.Success(user))
}

type InfoReq struct {
	Nickname string `json:"nickname"`
	Desc     string `json:"desc"`
	Avatar   string `json:"avatar"`
}

// SetUserInfo 设置个人中心信息
func (h *AuthHandler) SetUserInfo(ctx *gin.Context) {
	res := &result.Result{}
	req := &InfoReq{}

	token := ctx.Request.Header.Get("token")
	claims, _ := util.ParseToken(token)
	id := claims.Id

	err := ctx.BindJSON(req)
	if err != nil {
		log.Printf("json error => %s", err)
		return
	}

	err = model.SetUserInfo(id, req.Nickname, req.Desc, req.Avatar)
	if err != nil {
		log.Printf("mysql error => %s", err)
		ctx.JSON(200, res.Fail(4001, "更新失败！"))
		return
	}
	user, err := model.GetUserInfoByUserId(id)
	ctx.JSON(http.StatusOK, res.Success(user))
}
