package auth_service

import (
	"auth-service/util"
	authgrpc "commons/api/auth/gen"
	"context"
	"encoding/json"
)

type AuthService struct {
	authgrpc.UnimplementedTokenServiceServer
}

// VerifyToken 校验token
func (*AuthService) VerifyToken(ctx context.Context, req *authgrpc.Req) (*authgrpc.Res, error) {
	token := req.Token
	claims, err := util.ParseToken(token)
	if err != nil {
		return nil, err
	}
	c, _ := json.Marshal(claims)
	return &authgrpc.Res{
		Info: c,
	}, nil
}
