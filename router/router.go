package router

import (
	"rest_echo/api"
	"rest_echo/api/middlewares"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// router groups
	adminGroup := e.Group("/admin")
	jwtGroup := e.Group("/jwt")

	// set all middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetCompleteLogMiddlware(e)

	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetJwtMiddlewares(jwtGroup)

	// set main routes
	api.MainGroup(e)

	// set group routes
	api.AdminGroup(adminGroup)
	api.JwtGroup(jwtGroup)

	return e
}
