package api

import (
	"github.com/sepulsa/rest_echo/api/handlers"

	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.GET("/login", handlers.Login)
	e.GET("/yallo", handlers.Yallo)
	e.GET("/cats/:data", handlers.GetCats)

	e.POST("/cats", handlers.AddCat)
	e.POST("/dogs", handlers.AddDog)
	e.POST("/hamsters", handlers.AddHamster)

	e.GET("/users", handlers.GetUsers)
	e.GET("/user", handlers.GetUserById)
	e.POST("/users", handlers.AddUser)
	e.PUT("/users", handlers.EditUser)
	e.DELETE("/users", handlers.DeleteUser)
}
