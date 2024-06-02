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
	volunteerRepo := repositories.NewVolunteerRepository(database.DB)
	applicationRepo := repositories.NewApplicationRepository(database.DB)
	articleRepo := repositories.NewArticleRepository(database.DB)
	commentRepo := repositories.NewCommentRepository(database.DB)
	likesCommentRepo := repositories.NewLikesCommentRepository(database.DB)
	testimoniVolunteerRepo := repositories.NewTestimoniVolunteerRepository(database.DB)

	// Services
	userService := service.NewUserService(userRepo)
	volunteerService := service.NewVolunteerService(volunteerRepo)
	applicationService := service.NewApplicationService(applicationRepo)
	articleService := service.NewArticleService(articleRepo)
	commentService := service.NewCommentService(commentRepo)
	likesCommentService := service.NewLikesCommentService(likesCommentRepo)
	testimoniVolunteerService := service.NewTestimoniVolunteerService(testimoniVolunteerRepo)

	// Handlers
	userHandler := handler.NewUserHandler(userService)
	volunteerHandler := handler.NewVolunteerHandler(volunteerService)
	applicationHandler := handler.NewApplicationHandler(applicationService)
	articleHandler := handler.NewArticleHandler(articleService)
	commentHandler := handler.NewCommentHandler(commentService)
	likesCommentHandler := handler.NewLikesCommentHandler(likesCommentService)
	testimoniVolunteerHandler := handler.NewTestimoniVolunteerHandler(testimoniVolunteerService)

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
}
