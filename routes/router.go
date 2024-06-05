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

	// Repositories
	userRepo := repositories.NewUserRepository(database.DB)
	adminRepo := repositories.NewAdminRepository(database.DB)
	volunteerRepo := repositories.NewVolunteerRepository(database.DB)
	applicationRepo := repositories.NewApplicationRepository(database.DB)
	articleRepo := repositories.NewArticleRepository(database.DB)
	commentRepo := repositories.NewCommentRepository(database.DB)
	likesCommentRepo := repositories.NewLikesCommentRepository(database.DB)
	testimoniVolunteerRepo := repositories.NewTestimoniVolunteerRepository(database.DB)

	// Services
	fundraisingRepo := repositories.NewFundraisingRepository(database.DB)
	donationRepo := repositories.NewDonationRepository(database.DB)
	organizationRepo := repositories.NewOrganizationRepository(database.DB)

	userService := service.NewUserService(userRepo)
	adminService := service.NewAdminService(adminRepo)
	volunteerService := service.NewVolunteerService(volunteerRepo)
	applicationService := service.NewApplicationService(applicationRepo)
	articleService := service.NewArticleService(articleRepo)
	commentService := service.NewCommentService(commentRepo)
	likesCommentService := service.NewLikesCommentService(likesCommentRepo)
	testimoniVolunteerService := service.NewTestimoniVolunteerService(testimoniVolunteerRepo)

	// Handlers
	fundraisingService := service.NewFundraisingService(fundraisingRepo)
	donationService := service.NewDonationService(donationRepo, fundraisingRepo)
	organizationService := service.NewOrganizationService(organizationRepo)

	userHandler := handler.NewUserHandler(userService)
	adminHandler := handler.NewAdminHandler(adminService)
	volunteerHandler := handler.NewVolunteerHandler(volunteerService)
	applicationHandler := handler.NewApplicationHandler(applicationService)
	articleHandler := handler.NewArticleHandler(articleService)
	commentHandler := handler.NewCommentHandler(commentService)
	likesCommentHandler := handler.NewLikesCommentHandler(likesCommentService)
	testimoniVolunteerHandler := handler.NewTestimoniVolunteerHandler(testimoniVolunteerService)
	fundraisingHandler := handler.NewFundraisingHandler(fundraisingService, donationService)
	donationHandler := handler.NewDonationHandler(donationService, userService)
	organizatonHandler := handler.NewOrganizationHandler(organizationService)

	api := router.Group("api/v1")

	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.POST("/forget-password", userHandler.ForgetPassword)
	api.POST("/reset-password", userHandler.ResetPassword)

	api.POST("/transactions/notification", donationHandler.GetPaymentCallback)

	api.Use(jwt, routeMiddleware.UserMiddleware)

	api.GET("/organizations", organizatonHandler.GetOrganizations)
	api.GET("/organizations/:id", organizatonHandler.GetOrganizationByID)
	api.POST("/organizations", organizatonHandler.CreateOrganization)

	api.GET("/fundraisings", fundraisingHandler.GetFundraisings)
	api.GET("/fundraisings/top", fundraisingHandler.GetTopFundraisings)
	api.GET("/fundraising/:id", fundraisingHandler.GetFundraisingByID)
	api.GET("/fundraising-categories", fundraisingHandler.GetAllFundraisingCategories)
	api.GET("/fundraisings/:category_id", fundraisingHandler.GetFundraisingsByCategoryID)

	api.POST("/fundraisings", fundraisingHandler.CreateFundraisingContent)

	api.POST("/fundraising/:id/donations", donationHandler.CreateDonation)

	api.GET("/history/donations", donationHandler.GetUserDonation)
	api.GET("/history/donations/:id", donationHandler.GetDonationByID)

	api.POST("/comments/:comment_id/like", donationHandler.LikeComment)
	api.DELETE("/comments/:comment_id/unlike", donationHandler.UnLikeComment)

	api.GET("/home", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	api.POST("/volunteer/register", volunteerHandler.CreateVolunteer)
	api.GET("/volunteer/:id", volunteerHandler.GetVolunteerByID)
	api.GET("/volunteers", volunteerHandler.GetAllVolunteers)
	api.POST("/volunteer/:id/register", applicationHandler.RegisterApplication)

	// Article routes
	api.POST("/articles", articleHandler.CreateArticle)
	api.PUT("/articles/:id", articleHandler.UpdateArticle)
	api.GET("/articles/:id", articleHandler.GetArticleByID)
	api.GET("/articles", articleHandler.GetAllArticles)
	api.DELETE("/articles/:id", articleHandler.DeleteArticle)

	// Comment routes
	api.POST("/comments", commentHandler.CreateComment)
	api.PUT("/comments/:id", commentHandler.UpdateComment)
	api.GET("/comments/:id", commentHandler.GetCommentByID)
	api.GET("/comments", commentHandler.GetAllComments)
	api.DELETE("/comments/:id", commentHandler.DeleteComment)

	// LikesComment routes
	api.POST("/likes-comments", likesCommentHandler.CreateLikesComment)
	api.DELETE("/likes-comments/:id", likesCommentHandler.DeleteLikesComment)

	// TestimoniVolunteer routes
	api.POST("/testimoni-volunteers", testimoniVolunteerHandler.CreateTestimoniVolunteer)
	api.GET("/testimoni-volunteers/:id", testimoniVolunteerHandler.GetTestimoniVolunteerByID)
	api.GET("/testimoni-volunteers", testimoniVolunteerHandler.GetAllTestimoniVolunteers)
	api.DELETE("/testimoni-volunteers/:id", testimoniVolunteerHandler.DeleteTestimoniVolunteer)

	admin := router.Group("api/v1/admin")

	admin.POST("/login", adminHandler.Login)
	admin.Use(jwt, routeMiddleware.AdminMiddleware)
}
