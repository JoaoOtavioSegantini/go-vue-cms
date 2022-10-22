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
	router.Static("/uploads", "./uploads")

	router.Use(middlewares.CORSMiddleware())

	api := router.Group("/api")
	{
		api.POST("/token", controller.GenerateToken)
		api.POST("/user/register", controller.RegisterUser)
		api.POST("/v1/image/upload", func(ctx *gin.Context) {
			tok := ctx.Request.URL.Query().Get("token")
			url := os.Getenv("SERVER_URL")

			err := auth.ValidateToken(tok)
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				ctx.Abort()
				return
			}

			file, _ := ctx.FormFile("upload")
			err = ctx.SaveUploadedFile(file, "uploads/"+file.Filename)

			if err != nil {
				ctx.String(http.StatusInternalServerError, err.Error())
				ctx.Abort()
				return
			}

			ctx.JSON(http.StatusAccepted, gin.H{"token": tok, "upload": file.Filename, "uploaded": true, "url": url + file.Filename})
		})

	}
	// Simple group: v1
	v1 := router.Group("/api/v1").Use(middlewares.Auth())
	{
		v1.GET("/posts", postController.FindAll)
		v1.POST("/posts", postController.Save)
		v1.PUT("/posts/:id", postController.Update)
		v1.GET("/posts/:id", postController.FindById)
		v1.DELETE("/posts/:id", postController.Delete)
		v1.GET("/pages", pageController.FindAll)
		v1.POST("/pages", pageController.Save)
		v1.PUT("/pages/:id", pageController.Update)
		v1.GET("/pages/:id", pageController.FindById)
		v1.DELETE("/pages/:id", pageController.Delete)
		v1.GET("/users", userController.FindAll)
		v1.POST("/users", userController.Save)
		v1.PUT("/users/:id", userController.Update)
		v1.GET("/users/:id", userController.FindById)
		v1.DELETE("/users/:id", userController.Delete)
	}

	router.Run(":8000")
}
