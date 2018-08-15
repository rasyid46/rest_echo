package api

import (
	"rest_echo/api/handlers"

	"github.com/labstack/echo"
)

func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
