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

	routeMiddleware.LogMiddleware(router)

	userRepo := repositories.NewUserRepository(database.DB)
	volunteerRepo := repositories.NewVolunteerRepository(database.DB)
	applicationRepo := repositories.NewApplicationRepository(database.DB)

	userService := service.NewUserService(userRepo)
	volunteerService := service.NewVolunteerService(volunteerRepo)
	applicationService := service.NewApplicationService(applicationRepo)

	userHandler := handler.NewUserHandler(userService)
	volunteerHandler := handler.NewVolunteerHandler(volunteerService)
	applicationHandler := handler.NewApplicationHandler(applicationService)

	api := router.Group("api/v1")

	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.POST("/forget-password", userHandler.ForgetPassword)
	api.POST("/reset-password", userHandler.ResetPassword)
	api.Use(jwt, routeMiddleware.UserMiddleware)

	api.GET("/home", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	api.POST("/volunteer/register", volunteerHandler.CreateVolunteer)
	api.GET("/volunteer/:id", volunteerHandler.GetVolunteerByID)
	api.GET("/volunteers", volunteerHandler.GetAllVolunteers)

	api.POST("/volunteer/:id/register", applicationHandler.RegisterApplication)
}
