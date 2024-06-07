package routes

import (
	"capstone/handler"
	routeMiddleware "capstone/middlewares"
	"capstone/repositories"
	"capstone/service"
	"capstone/utils/database"
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(router *echo.Echo) {
	var jwt = echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	routeMiddleware.LogMiddleware(router)

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

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
	adminService := service.NewAdminService(adminRepo, userRepo)
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
	adminHandler := handler.NewAdminHandler(adminService, volunteerService)
	volunteerHandler := handler.NewVolunteerHandler(volunteerService, applicationService)
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

	api.GET("/fundraisings", fundraisingHandler.GetFundraisings)
	api.GET("/fundraisings/top", fundraisingHandler.GetTopFundraisings)
	api.GET("/fundraising/:id", fundraisingHandler.GetFundraisingByID)
	api.GET("/fundraising-categories", fundraisingHandler.GetAllFundraisingCategories)
	api.GET("/fundraisings/categories/:category_id", fundraisingHandler.GetFundraisingsByCategoryID)

	api.POST("/fundraising/:id/donations", donationHandler.CreateDonation)

	api.GET("/history/donations", donationHandler.GetUserDonation)
	api.GET("/history/donations/:id", donationHandler.GetDonationByID)

	api.POST("/comments/:comment_id/like", donationHandler.LikeComment)
	api.DELETE("/comments/:comment_id/unlike", donationHandler.UnLikeComment)

	// Volunteer
	api.GET("/volunteer/:id", volunteerHandler.GetVolunteerByID)
	api.GET("/volunteers", volunteerHandler.GetAllVolunteers)
	api.GET("/volunteer/:id/confirm", volunteerHandler.ConfirmVolunteer)
	api.POST("/volunteer/:id/apply", volunteerHandler.ApplyForVolunteer)

	// Application routes
	api.POST("/volunteer/:id/register", applicationHandler.RegisterApplication)
	api.GET("/volunteer/:id/applications", applicationHandler.GetAllApplications)
	api.GET("/volunteer/applications/:id", applicationHandler.GetApplicationByID)
	api.DELETE("/volunteer/applications/:id", applicationHandler.DeleteApplicationByID)

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
	api.GET("/likes-comments/:id", likesCommentHandler.GetLikesCommentByID)
	api.GET("/likes-comments", likesCommentHandler.GetAllLikesComments)

	// TestimoniVolunteer routes
	api.POST("/volunteer/:id/testimoni-volunteers", testimoniVolunteerHandler.CreateTestimoniVolunteer)
	api.GET("/testimoni-volunteers/:id", testimoniVolunteerHandler.GetTestimoniVolunteerByID)
	api.GET("/testimoni-volunteers", testimoniVolunteerHandler.GetAllTestimoniVolunteers)
	api.DELETE("/testimoni-volunteers/:id", testimoniVolunteerHandler.DeleteTestimoniVolunteer)

	// Admin
	admin := router.Group("api/v1/admin")

	admin.POST("/login", adminHandler.Login)
	admin.Use(jwt, routeMiddleware.AdminMiddleware)

	admin.GET("/users", adminHandler.GetAllUsers)
	admin.GET("/users/:id/donations", adminHandler.GetDetailUserWithDonations)
	admin.DELETE("/users/:id", adminHandler.DeleteUser)

	admin.GET("/fundraisings", adminHandler.GetFundraisings)
	admin.POST("/fundraisings", adminHandler.CreateFundraisingContent)
	admin.GET("/fundraisings/:id", adminHandler.GetDetailFundraising)
	admin.GET("/fundraisings/:id/donations", adminHandler.GetDonationsByFundraisingID)
	admin.DELETE("/fundraisings/:id", adminHandler.DeleteFundraising)
	admin.PUT("/fundraisings/:id", adminHandler.EditFundraising)

	admin.POST("/organizations", organizatonHandler.CreateOrganization)
	admin.GET("/organizations", adminHandler.GetAllOrganizations)
	admin.PUT("/organizations/:id", adminHandler.EditOrganization)
	admin.DELETE("/organizations/:id", adminHandler.DeleteOrganization)

	admin.GET("/volunteers", volunteerHandler.GetAllVolunteers)
	admin.POST("/volunteers", volunteerHandler.CreateVolunteer)
	admin.PUT("/volunteers/:id", volunteerHandler.UpdateVolunteer)
	admin.DELETE("/volunteers/:id", volunteerHandler.DeleteVolunteer)
}
