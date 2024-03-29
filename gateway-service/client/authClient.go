package client

import (
	"commons/config"
	"commons/result"
	"context"
	"github.com/carlmjohnson/requests"
	"log"
)

// VerifyToken 校验token
func VerifyToken(ctx context.Context, token string) (*result.Result, error) {
	r := &result.Result{}
	err := requests.URL(config.Conf.SEVERURL.AuthHttp).Path("/verify").BodyJSON(token).ToJSON(&r).Fetch(ctx)
	if err != nil {
		log.Printf("err => %s", err)
		return nil, err
	}
	log.Printf("result => %+v", r)
	return r, nil
}
