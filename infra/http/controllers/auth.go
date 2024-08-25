package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	"smart-mess/domain"
)

type AuthController struct {
	baseCtx context.Context
	authuc  domain.IAuthService
}

func NewAuthController(
	baseCtx context.Context,
	authuc domain.IAuthService,
) *AuthController {
	return &AuthController{
		baseCtx: baseCtx,
		authuc:  authuc,
	}
}
func (ctrlr *AuthController) Login(c echo.Context) error {
	//ctx := domain.ContextWithValue(ctrlr.baseCtx, consts.ContextKeyUser, parseUser(c))
	//var err error
	//var cred *types.LoginReq
	//if err := c.Bind(&cred); err != nil {
	//	return c.JSON(http.StatusBadRequest, &types.ValidationError{
	//		Error: err.Error(),
	//	})
	//}
	//
	//var res *types.LoginResp
	//
	//if res, err = authsvc.Login(cred); err != nil {
	//	switch {
	//	case errors.Is(err, errutil.ErrInvalidEmail), errors.Is(err, errutil.ErrInvalidPassword):
	//		return c.JSON(http.StatusUnauthorized, msgCtx.InvalidUserPassword())
	//	case errors.Is(err, errutil.ErrNotAdmin):
	//		return c.JSON(http.StatusForbidden, msgCtx.AccessForbidden())
	//	case errors.Is(err, errutil.ErrCreateJwt):
	//		return c.JSON(http.StatusInternalServerError, msgCtx.SomethingWentWrong().WithValidationError(errutil.ErrCreateJwt))
	//	default:
	//		return c.JSON(http.StatusInternalServerError, msgCtx.SomethingWentWrong())
	//	}
	//}
	//
	//return c.JSON(http.StatusOK, res)
	return nil
}
