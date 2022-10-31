package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joaotavioos/cms-server/auth"
	"github.com/joaotavioos/cms-server/controller"
	"github.com/joaotavioos/cms-server/database"
	"github.com/joaotavioos/cms-server/middlewares"
	"github.com/joaotavioos/cms-server/service"
	"github.com/joho/godotenv"
)

var (
	postService    service.PostService       = service.New()
	postController controller.PostController = controller.New(postService)
	pageService    service.PageService       = service.NewPage()
	pageController controller.PageController = controller.NewPage(pageService)
	userService    service.UserService       = service.NewUsers()
	userController controller.UserController = controller.NewUsers(userService)
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {

	database.Connect(os.Getenv("DATABASE_URL"))
	database.Migrate()

	router := gin.Default()
	router.Use(middlewares.CORS())

	router.Static("/uploads", "./uploads")
	router.LoadHTMLGlob("templates/*.html")

	api := router.Group("/api")
	{
		api.POST("/token", controller.GenerateToken)
		api.POST("/user/register", controller.RegisterUser)
		api.POST("/v1/image/upload", func(ctx *gin.Context) {
			tok := ctx.Request.URL.Query().Get("token")
			url := os.Getenv("SERVER_URL")

			err := auth.ValidateToken(tok)
			if err != nil {
				ctx.Abort()
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}

			file, _ := ctx.FormFile("upload")
			err = ctx.SaveUploadedFile(file, "uploads/"+file.Filename)

			if err != nil {
				ctx.Abort()
				ctx.String(http.StatusInternalServerError, err.Error())
				return
			}

			ctx.JSON(http.StatusAccepted, gin.H{"token": tok, "upload": file.Filename, "uploaded": true, "url": url + file.Filename})
		})

	}

	pages_site := router.Group("/home")
	{
		pages_site.GET("", controller.Home)
		pages_site.GET("/:slug", controller.View)
		pages_site.GET("/articles/:slug", controller.Show)

	}

	router.Use(middlewares.Auth())
	// Simple group: v1
	v1 := router.Group("/api/v1/site-admin-posts")
	{
		v1.GET("", postController.FindAll)
		v1.POST("", postController.Save)
		v1.PUT("/:id", postController.Update)
		v1.GET("/:id", postController.FindById)
		v1.DELETE("/:id", postController.Delete)

	}

	pgs := router.Group("api/v1/res-data")
	{
		pgs.GET("", pageController.FindAll)
		pgs.POST("", pageController.Save)
		pgs.PUT("/:id", pageController.Update)
		pgs.GET("/:id", pageController.FindById)
		pgs.DELETE("/:id", pageController.Delete)

	}

	u := router.Group("api/v1/users-site-admin")
	{
		u.GET("", userController.FindAll)
		u.POST("", userController.Save)
		u.PUT("/:id", userController.Update)
		u.GET("/:id", userController.FindById)
		u.DELETE("/:id", userController.Delete)
	}

	router.Run(":8000")
}
