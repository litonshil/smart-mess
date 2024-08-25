package auth

import (
	"smart-mess/domain"
	"smart-mess/types"
)

type AuthService struct {
	repo domain.IAuthRepo
}

func NewAuthService(authRepo domain.IAuthRepo) *AuthService {
	return &AuthService{
		repo: authRepo,
	}
}

func (service *AuthService) Login(req *types.LoginReq) (*types.LoginResp, error) {
	_ = req
	//if err := service.repo.SignUp(user); err != nil {
	//	return errors.New("User was not created")
	//}

	return nil, nil
}
