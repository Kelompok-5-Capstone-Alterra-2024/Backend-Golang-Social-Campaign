package routes

import (
	"capstone/handler"
	routeMiddleware "capstone/middlewares"
	"capstone/repositories"
	"capstone/service"
	"capstone/utils/database"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt"

	"github.com/labstack/echo/v4"
)

func NewRouter(router *echo.Echo) {
	var jwt = echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	// e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	routeMiddleware.LogMiddleware(router)

	userRepo := repositories.NewUserRepository(database.DB)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	api := router.Group("api/v1")

	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.POST("/forget-password", userHandler.ForgetPassword)
	api.POST("/reset-password", userHandler.ResetPassword)

	api.GET("/profile", userHandler.GetUserProfile)
	api.PUT("/profile/edit", userHandler.EditProfile)
	api.PUT("/profile/change-password", userHandler.ChangePassword)

	api.GET("/profile/history/volunteers", userHandler.GetHistoryVolunteer)
	api.GET("/profile/history/volunteers/:id", userHandler.GetHistoryVolunteerDetail)

	api.Use(jwt, routeMiddleware.UserMiddleware)

	api.GET("/home", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

}
