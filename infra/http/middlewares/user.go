package middlewares

import (
	"github.com/labstack/echo/v4"
	"smart-mess/config"
	"smart-mess/domain"
)

func GenerateMetadata(c echo.Context, user *domain.User) *domain.User {
	if user == nil {
		user = &domain.User{}
	}

	var body interface{}
	_ = BindBody(c, &body)
	appkey := c.Request().Header.Get(config.App().AppKeyHeader)
	if appkey != "" {
		appkey = "internal call (app key provided)"
	}
	//serviceName := c.Request().Header.Get(headerShadowchefServiceName)
	// metadata will be passed as slack logger metadata
	metadata := domain.Meta{
		Method: c.Request().Method,
		URI:    c.Request().RequestURI,
		//ServiceName: &serviceName,
		AppKey: &appkey,
		Profile: domain.Profile{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		},
		//Payload: body,
		//
		//Profile: domain.Profile{
		//	FirstName: user.FirstName,
		//	LastName:  user.LastName,
		//	Email:     user.Email,
		//},
		//Payload: body,

		//Payload: body,
		//
		//Profile: domain.Profile{
		//	FirstName: user.FirstName,
		//	LastName:  user.LastName,
		//	Email:     user.Email,
		//},
		//Payload: body,

		//Payload: body,

		//test 2
		//test 2
		//test 2
		//test 2

		//AppKey: &appkey,
		//AppKey: &appkey,
		//AppKey: &appkey,

		//test
		//test
		//test
		//AppKey: &appkey,
		//AppKey: &appkey,
		//AppKey: &appkey,
	}
	user.Metadata = metadata
	return user
}
