package routes

import (
	"capstone/handler"
	"capstone/repositories"
	"capstone/service"
	"capstone/utils/database"

	"github.com/labstack/echo/v4"
)

func NewRouter(router *echo.Echo) {

	// e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	userRepo := repositories.NewUserRepository(database.DB)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	api := router.Group("api/v1")

	api.POST("/register", userHandler.Register)

}
