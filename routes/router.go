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
	fundraisingRepo := repositories.NewFundraisingRepository(database.DB)
	donationRepo := repositories.NewDonationRepository(database.DB)
	organizationRepo := repositories.NewOrganizationRepository(database.DB)

	userService := service.NewUserService(userRepo)
	fundraisingService := service.NewFundraisingService(fundraisingRepo)
	donationService := service.NewDonationService(donationRepo, fundraisingRepo)
	organizationService := service.NewOrganizationService(organizationRepo)

	userHandler := handler.NewUserHandler(userService)
	fundraisingHandler := handler.NewFundraisingHandler(fundraisingService, donationService)
	donationHandler := handler.NewDonationHandler(donationService, userService)
	organizatonHandler := handler.NewOrganizationHandler(organizationService)

	api := router.Group("api/v1")

	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.POST("/forget-password", userHandler.ForgetPassword)
	api.POST("/reset-password", userHandler.ResetPassword)

	api.GET("/organizations", organizatonHandler.GetOrganizations)
	api.GET("/organizations/:id", organizatonHandler.GetOrganizationByID)

	api.GET("/fundraisings", fundraisingHandler.GetFundraisings)
	api.GET("/fundraising/:id", fundraisingHandler.GetFundraisingByID)
	api.GET("/fundraising-categories", fundraisingHandler.GetAllFundraisingCategories)
	api.GET("/fundraisings/:category_id", fundraisingHandler.GetFundraisingsByCategoryID)
	api.POST("/transactions/notification", donationHandler.GetPaymentCallback)

	api.Use(jwt, routeMiddleware.UserMiddleware)

	api.POST("/fundraising/:id/donations", donationHandler.CreateDonation)

	api.GET("/history/donations", donationHandler.GetUserDonation)
	api.GET("/history/donations/:id", donationHandler.GetDonationByID)

	api.POST("/comments/:comment_id/like", donationHandler.LikeComment)
	api.DELETE("/comments/:comment_id/unlike", donationHandler.UnLikeComment)

	api.GET("/home", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

}
