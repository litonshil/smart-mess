package user

import (
	"context"
	"smart-mess/domain"
	"smart-mess/types"
)

type UserService struct {
	repo domain.UserRepo
}

func NewUserService(userRepo domain.UserRepo) *UserService {
	return &UserService{
		repo: userRepo,
	}
}

func (service *UserService) CreateUser(ctx context.Context, req types.UserReq) (*types.UserResp, error) {
	_ = req
	_ = ctx
	//if err := service.repo.SignUp(user); err != nil {
	//	return errors.New("User was not created")
	//}

	return nil, nil
}
