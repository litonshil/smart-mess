package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	"smart-mess/domain"
	"smart-mess/utils/consts"
)

type UserController struct {
	baseCtx context.Context
	useruc  domain.UserUseCase
}

func NewUserController(
	baseCtx context.Context,
	useruc domain.UserUseCase,
) *UserController {
	return &UserController{
		baseCtx: baseCtx,
		useruc:  useruc,
	}
}
func (cntlr *UserController) CreateUser(c echo.Context) error {
	ctx := domain.ContextWithValue(cntlr.baseCtx, consts.ContextKeyUser, parseUser(c))
	_ = ctx
	//var err error

	return nil
	//return c.JSON(http.StatusOK, res)
}
