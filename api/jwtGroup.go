package api

import (
	"github.com/sepulsa/rest_echo/api/handlers"

	"github.com/labstack/echo"
)

func JwtGroup(g *echo.Group) {
	g.GET("/main", handlers.MainJwt)
}
