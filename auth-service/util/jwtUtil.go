package util

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

// MyClaims 生成和校验jwt token
type MyClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// TokenExpireDuration token过期时间 7天
const TokenExpireDuration = time.Hour * 24 * 7

var MySecret = []byte("DataPulse")

// CreateToken 生成JwtToken
func CreateToken(id int, username string) (string, error) {
	c := MyClaims{
		Id:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), //有效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                          // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                          // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString(MySecret)
	return tokenString, err
}

// ParseToken 校验JwtToken
func ParseToken(tokenStr string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		log.Printf("err => %s", err)
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
