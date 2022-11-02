package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joaotavioos/cms-server/database"
	"github.com/joaotavioos/cms-server/routes"
	"github.com/joho/godotenv"
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
	r := gin.Default()

	router := routes.SetupRouter(r, "templates/*.html")
	router.Run(":8000")
}
