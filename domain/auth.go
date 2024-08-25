package domain

import "smart-mess/types"

type IAuthService interface {
	Login(req *types.LoginReq) (*types.LoginResp, error)
}

type IAuthRepo interface {
	Login(req *types.LoginReq) (*types.LoginResp, error)
}
