package db

import (
	"smart-mess/domain"
	"smart-mess/types"
)

func (repo *Repository) CreateUser(user domain.User) (*types.UserResp, error) {
	//if err := repo.db.Create(user).Error; err != nil {
	//	return err
	//}

	return nil, nil
}
