package routes

import (
	"github.com/labstack/echo/v4"
	c "smart-mess/infra/http/controllers"
)

type Routes struct {
	echo           *echo.Echo
	authController *c.AuthController
	userController *c.UserController
}

func New(
	e *echo.Echo,
	authController *c.AuthController,
	userController *c.UserController,
) *Routes {
	return &Routes{
		echo:           e,
		authController: authController,
		userController: userController,
	}
}

func (r *Routes) Init() {
	e := r.echo
	//middlewares.Init(e)

	//e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	g := e.Group("/v1")
	g.POST("/user", r.userController.CreateUser)
	g.POST("/login", r.authController.Login)
	//g.POST("/v1/login", ac.Login)
	//g.POST("/v1/logout", ac.Logout)
	//g.POST("/v1/token/refresh", ac.RefreshToken)
	//g.GET("/v1/token/verify", ac.VerifyToken)

}
