package main

import (
	"belajar/config"
	"belajar/docs"
	"belajar/routes"
	"belajar/utils"
	"log"

	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
	environment := utils.Getenv("ENVIRONMENT", "development")
	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading env file")
		}
	}

	docs.SwaggerInfo.Title = "API Belajar Shop"
	docs.SwaggerInfo.Description = "Ini belajar API backend"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Version = "0.5"

	db := config.ConnectDatabase()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	r := routes.SetupRouter(db)
	r.Run("localhost:8080")
}