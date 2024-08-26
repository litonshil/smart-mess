package domain

import (
	"context"
	"smart-mess/types"
)

type User struct {
	ID       int  `json:"id"`
	Metadata Meta `json:"meta"`
	Profile
}

type Profile struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type Meta struct {
	Method      string  `json:"method"`
	URI         string  `json:"uri"`
	ServiceName *string `json:"serviceName,omitempty"`
	AppKey      *string `json:"app-key,omitempty"`
	Profile
	Payload interface{} `json:"payload"`
}

type UserUseCase interface {
	CreateUser(ctx context.Context, req types.UserReq) (*types.UserResp, error)
}

type UserRepo interface {
	CreateUser(req User) (*types.UserResp, error)
}
