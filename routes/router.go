package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(router *echo.Echo) {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

}
