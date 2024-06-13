package main

import (
	"capstone/routes"
	"capstone/utils/database"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Failed to fetch .env file")
		}
	}

	database.InitDB()
	// migration.Migration()

	app := echo.New()
	routes.NewRouter(app)

	app.Start(":8080")
}
