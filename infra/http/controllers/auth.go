package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"smart-mess/domain"
	"smart-mess/types"
	"smart-mess/utils/consts"
	"smart-mess/utils/errutil"
	"smart-mess/utils/msgutil"
)

type AuthController struct {
	baseCtx context.Context
	authuc  domain.AuthUseCase
}

func NewAuthController(
	baseCtx context.Context,
	authuc domain.AuthUseCase,
) *AuthController {
	return &AuthController{
		baseCtx: baseCtx,
		authuc:  authuc,
	}
}
func (ctrlr *AuthController) Login(c echo.Context) error {
	ctx := domain.ContextWithValue(ctrlr.baseCtx, consts.ContextKeyUser, parseUser(c))
	var err error
	var cred *types.LoginReq
	if err := c.Bind(&cred); err != nil {
		return c.JSON(http.StatusBadRequest, &types.ValidationError{
			Error: err.Error(),
		})
	}

	var res *types.LoginResp

	if res, err = ctrlr.authuc.Login(ctx, cred); err != nil {
		switch err {
		case errutil.ErrInvalidEmail, errutil.ErrInvalidPassword, errutil.ErrNotAdmin:
			unAuthErr := msgutil.InvalidCredentialsMsg()
			return c.JSON(http.StatusForbidden, unAuthErr)
		case errutil.ErrCreateJwt:
			serverErr := errutil.NewError("failed to create jwt token")
			return c.JSON(http.StatusInternalServerError, serverErr)
		case errutil.ErrStoreTokenUuid:
			serverErr := errutil.NewError("failed to store jwt token uuid")
			return c.JSON(http.StatusInternalServerError, serverErr)
		default:
			serverErr := errutil.ErrSomethingWentWrong
			return c.JSON(http.StatusInternalServerError, serverErr)
		}
	}

	return c.JSON(http.StatusOK, res)
}
