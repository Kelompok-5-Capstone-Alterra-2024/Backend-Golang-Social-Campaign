package routes

import (
	"capstone/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() echo.Echo {
	database.InitDatabase()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return *e
}
