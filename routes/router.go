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

	// User repository, service, and handler
	userRepo := repositories.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Volunteer repository, service, and handler
	volunteerRepo := repositories.NewVolunteerRepository(database.DB)
	volunteerService := service.NewVolunteerService(volunteerRepo)
	volunteerHandler := handler.NewVolunteerHandler(volunteerService)

	// Application repository, service, and handler
	applicationRepo := repositories.NewApplicationRepository(database.DB)
	applicationService := service.NewApplicationService(applicationRepo)
	applicationHandler := handler.NewApplicationHandler(applicationService)

	api := router.Group("api/v1")

	// User routes
	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.POST("/forget-password", userHandler.ForgetPassword)
	api.POST("/reset-password", userHandler.ResetPassword)

	api.Use(jwt, routeMiddleware.UserMiddleware)

	api.GET("/home", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	// Volunteer routes
	api.GET("/volunteer/:id", volunteerHandler.GetVolunteerByID)
	api.GET("/volunteers", volunteerHandler.GetAllVolunteer)
	api.POST("/volunteer/:id/register", applicationHandler.RegisterApplication)
	api.POST("/volunteer", volunteerHandler.CreateVolunteer)
}
