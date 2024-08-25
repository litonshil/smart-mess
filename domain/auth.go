package domain

import (
	"context"
	"smart-mess/types"
)

type AuthUseCase interface {
	Login(ctx context.Context, req *types.LoginReq) (*types.LoginResp, error)
}

type AuthRepo interface {
	Login(req *types.LoginReq) (*types.LoginResp, error)
}
