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

	// router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderAccessControlAllowOrigin},
	// }))
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// router.Use(middleware.CORS())

	// Repositories
	userRepo := repositories.NewUserRepository(database.DB)
	adminRepo := repositories.NewAdminRepository(database.DB)
	volunteerRepo := repositories.NewVolunteerRepository(database.DB)
	applicationRepo := repositories.NewApplicationRepository(database.DB)
	articleRepo := repositories.NewArticleRepository(database.DB)
	commentRepo := repositories.NewCommentRepository(database.DB)
	likesCommentRepo := repositories.NewLikesCommentRepository(database.DB)
	testimoniVolunteerRepo := repositories.NewTestimoniVolunteerRepository(database.DB)

	fundraisingRepo := repositories.NewFundraisingRepository(database.DB)
	donationRepo := repositories.NewDonationRepository(database.DB)
	donationManualRepo := repositories.NewDonationManualRepository(database.DB)
	organizationRepo := repositories.NewOrganizationRepository(database.DB)
	transactionRepo := repositories.NewTransactionRepository(database.DB)

	userService := service.NewUserService(userRepo)
	adminService := service.NewAdminService(adminRepo, userRepo)
	volunteerService := service.NewVolunteerService(volunteerRepo)
	applicationService := service.NewApplicationService(applicationRepo)
	articleService := service.NewArticleService(articleRepo)
	commentService := service.NewCommentService(commentRepo)
	likesCommentService := service.NewLikesCommentService(likesCommentRepo)
	testimoniVolunteerService := service.NewTestimoniVolunteerService(testimoniVolunteerRepo)

	fundraisingService := service.NewFundraisingService(fundraisingRepo)
	donationService := service.NewDonationService(donationRepo, fundraisingRepo)
	donationManualService := service.NewDonationManualService(donationManualRepo, fundraisingRepo)
	organizationService := service.NewOrganizationService(organizationRepo)
	transactionService := service.NewTransactionService(transactionRepo, adminRepo)

	userHandler := handler.NewUserHandler(userService, fundraisingService)
	adminHandler := handler.NewAdminHandler(adminService, volunteerService, articleService, commentService)
	volunteerHandler := handler.NewVolunteerHandler(volunteerService, applicationService, testimoniVolunteerService)
	applicationHandler := handler.NewApplicationHandler(applicationService)
	articleHandler := handler.NewArticleHandler(articleService)
	commentHandler := handler.NewCommentHandler(commentService)
	likesCommentHandler := handler.NewLikesCommentHandler(likesCommentService)
	testimoniVolunteerHandler := handler.NewTestimoniVolunteerHandler(testimoniVolunteerService)
	fundraisingHandler := handler.NewFundraisingHandler(fundraisingService, donationManualService)
	donationHandler := handler.NewDonationHandler(donationService, userService)
	donationManualHandler := handler.NewDonationManualHandler(donationManualService, userService, fundraisingService)
	organizatonHandler := handler.NewOrganizationHandler(organizationService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	api := router.Group("/api/v1")

	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.POST("/forget-password", userHandler.ForgetPassword)
	api.POST("/reset-password", userHandler.ResetPassword)

	api.POST("/refresh-token", userHandler.RefreshToken)

	api.Use(jwt, routeMiddleware.UserMiddleware)

	api.GET("/notifications", userHandler.GetNotificationFundraising)

	api.GET("/profile", userHandler.GetUserProfile)
	api.PUT("/profile/edit", userHandler.EditProfile)
	api.PUT("/profile/change-password", userHandler.ChangePassword)

	api.GET("/profile/history/fundraisings", userHandler.GetHistoryDonation)
	api.GET("/history/volunteers", userHandler.GetHistoryVolunteer)
	api.GET("/history/volunteers/:id", userHandler.GetHistoryVolunteerDetail)
	api.POST("/transactions/notification", donationHandler.GetPaymentCallback)

	api.GET("/profile/bookmark/fundraisings", userHandler.GetBookmarkFundraising)
	api.POST("/fundraisings/bookmark/:id", userHandler.CreateBookmarkFundraising)
	api.DELETE("/fundraisings/bookmark/:id", userHandler.DeleteBookmarkFundraising)

	api.GET("/profile/bookmark/articles", userHandler.GetBookmarkArticle)
	api.POST("/articles/bookmark/:id", userHandler.CreateBookmarkArticle)
	api.DELETE("/articles/bookmark/:id", userHandler.DeleteBookmarkArticle)

	api.GET("/profile/bookmark/volunteers", userHandler.GetUserBookmarkVolunteer)
	api.POST("/volunteers/bookmark/:id", userHandler.CreateBookmarkVolunteer)
	api.DELETE("/volunteers/bookmark/:id", userHandler.DeleteBookmarkVolunteer)

	api.GET("/organizations", organizatonHandler.GetOrganizations)
	api.GET("/fundraising/organizations/:id", organizatonHandler.GetFundraisingsOrganizationByID)
	api.GET("/volunteer/organizations/:id", organizatonHandler.GetVolunteersByOrganizationByID)

	api.GET("/fundraisings", fundraisingHandler.GetFundraisings)
	api.GET("/fundraisings/top", fundraisingHandler.GetTopFundraisings)
	api.GET("/fundraising/:id", fundraisingHandler.GetFundraisingByID)
	api.GET("/fundraising-categories", fundraisingHandler.GetAllFundraisingCategories)
	api.GET("/fundraisings/categories/:category_id", fundraisingHandler.GetFundraisingsByCategoryID)

	api.POST("/fundraising/:id/donations", donationHandler.CreateDonation)

	api.POST("/fundraising/:id/donations-manual", donationManualHandler.CreateManualDonation)
	api.GET("/history/donations-manual", donationManualHandler.GetDonationManualByUserID)
	api.GET("/history/donations-manual/:id", donationManualHandler.GetDonationManualByID)

	api.POST("/donations-manual/comments/:id/like", donationManualHandler.LikeComment)
	api.DELETE("/donations-manual/comments/:id/unlike", donationManualHandler.UnlikeComment)

	api.GET("/history/donations", donationHandler.GetUserDonation)
	api.GET("/history/donations/:id", donationHandler.GetDonationByID)

	api.POST("/comments/:comment_id/like", donationHandler.LikeComment)
	api.DELETE("/comments/:comment_id/unlike", donationHandler.UnLikeComment)

	// Volunteer
	api.GET("/volunteer/:id", volunteerHandler.GetVolunteerByID)
	api.GET("/volunteers", volunteerHandler.GetAllVolunteers)
	api.GET("/volunteers/top", volunteerHandler.GetTopVolunteer)
	api.GET("/volunteer/:id/confirm", volunteerHandler.ConfirmVolunteer)
	api.GET("/volunteer/:id/apply", volunteerHandler.ApplyForVolunteer)

	// Application routes
	api.POST("/volunteer/:id/register", applicationHandler.RegisterApplication)
	api.GET("/volunteer/:id/applications", applicationHandler.GetAllApplications)
	api.GET("/volunteer/applications/:id", applicationHandler.GetApplicationByID)
	api.DELETE("/volunteer/applications/:id", applicationHandler.DeleteApplicationByID)

	// Article routes
	api.GET("/articles/:id", articleHandler.GetArticleByID)
	api.GET("/articles", articleHandler.GetAllArticles)
	api.GET("/articles/top", articleHandler.GetTopArticles)

	// Comment routes
	api.POST("/articles/:id/comments", commentHandler.CreateComment)
	api.GET("/articles/:id/comments", commentHandler.GetCommentsByArticleID)
	// api.PUT("/comments/:id", commentHandler.UpdateComment)
	// api.GET("/comments/:id", commentHandler.GetCommentByID)
	// api.GET("/comments", commentHandler.GetAllComments)
	// api.DELETE("/comments/:id", commentHandler.DeleteComment)

	// LikesComment routes
	api.POST("/article/comments/:id/like", likesCommentHandler.CreateLikesComment)
	api.DELETE("/article/comments/:id/unlike", likesCommentHandler.DeleteLikesComment)
	// api.GET("/likes-comments/:id", likesCommentHandler.GetLikesCommentByID)
	// api.GET("/likes-comments", likesCommentHandler.GetAllLikesComments)

	// TestimoniVolunteer routes
	api.POST("/volunteer/:id/testimoni-volunteers", testimoniVolunteerHandler.CreateTestimoniVolunteer)
	api.GET("/testimoni-volunteers/:id", testimoniVolunteerHandler.GetTestimoniVolunteerByID)
	api.GET("/testimoni-volunteers", testimoniVolunteerHandler.GetAllTestimoniVolunteers)
	api.GET("/volunteer/:id/testimoni-volunteers", testimoniVolunteerHandler.GetAllTestimoniVolunteersByVacancyID)
	api.DELETE("/testimoni-volunteers/:id", testimoniVolunteerHandler.DeleteTestimoniVolunteer)

	// Admin
	admin := router.Group("api/v1/admin")
	admin.POST("/refresh-token", adminHandler.RefreshTokenAdmin)

	admin.POST("/login", adminHandler.Login)
	admin.Use(jwt, routeMiddleware.AdminMiddleware)

	admin.GET("/users", adminHandler.GetAllUsers)
	admin.GET("/users/:id/donations", adminHandler.GetUserDonations)
	admin.GET("/users/:id/volunteers", adminHandler.GetUserVolunteers)
	admin.GET("/users/:id", adminHandler.GetUserDetail)
	admin.PUT("/users/:id", adminHandler.EditUsers)
	admin.DELETE("/users/:id", adminHandler.DeleteUser)

	admin.GET("/fundraisings", adminHandler.GetFundraisings)
	admin.POST("/fundraisings", adminHandler.CreateFundraisingContent)
	admin.GET("/fundraisings/:id", adminHandler.GetDetailFundraising)
	admin.GET("/fundraisings/:id/donations", adminHandler.GetDonationsByFundraisingID)
	admin.DELETE("/fundraisings/:id", adminHandler.DeleteFundraising)
	admin.PUT("/fundraisings/:id", adminHandler.EditFundraising)
	admin.GET("/fundraising-categories", fundraisingHandler.GetAllFundraisingCategories)

	admin.GET("/donations", adminHandler.GetAllDonationManual)
	admin.POST("/donations/:id", adminHandler.InputAmountDonationManual)

	admin.POST("/distributions", transactionHandler.CreateTransaction)
	admin.GET("/transactions-history", transactionHandler.GetTransactions)
	admin.GET("/transactions/:id", transactionHandler.GetTransactionByID)

	admin.POST("/organizations", organizatonHandler.CreateOrganization)
	admin.GET("/organizations", adminHandler.GetAllOrganizations)
	admin.GET("/organizations/:id", adminHandler.GetOrganizationByID)
	admin.PUT("/organizations/:id", adminHandler.EditOrganization)
	admin.DELETE("/organizations/:id", adminHandler.DeleteOrganization)
	admin.GET("/organizations/:id/fundraising", adminHandler.GetFundraisingByOrganization)
	admin.GET("/organizations/:id/volunteers", adminHandler.GetVolunteerByOrganization)

	admin.GET("/volunteers", adminHandler.GetAdminAllVolunteers)
	admin.GET("/volunteer/:id", volunteerHandler.GetVolunteerByID)
	admin.POST("/volunteers", volunteerHandler.CreateVolunteer)
	admin.PUT("/volunteers/:id", volunteerHandler.UpdateVolunteer)
	admin.DELETE("/volunteers/:id", volunteerHandler.DeleteVolunteer)

	admin.GET("/volunteers/:id/applications", volunteerHandler.GetAllApplyVolunteers)

	admin.GET("/articles", adminHandler.GetAdminAllArticle)
	admin.GET("/articles/:id", adminHandler.GetAdminArticleByID)
	admin.POST("/articles", articleHandler.CreateArticle)
	admin.PUT("/articles/:id", articleHandler.UpdateArticle)
	admin.DELETE("/articles/:id", articleHandler.DeleteArticle)
	admin.GET("/articles/:id/comments", commentHandler.GetCommentsByArticleID)

	admin.GET("/data-total-content", adminHandler.GetDataTotalContent)
	admin.GET("/transactions-daily", adminHandler.GetDailyDonationSummary)
	admin.GET("/articles-top", adminHandler.GetArticlesOrderedByBookmarks)
	admin.GET("/volunteers-top", volunteerHandler.GetTopVolunteer)
	admin.GET("/categories-top", adminHandler.GetCategoriesWithCount)
}
