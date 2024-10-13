package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"task-manager-api/db"
	"task-manager-api/migrations"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"task-manager-api/controllers"
	"task-manager-api/utils/config"
)

func main() {
	db.Connect()
	defer db.CloseConnection()

	// Run migrations
	migrations.Migrate()

	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	config.LoadEnv()

	// Middleware
	app.Use(logger.New())
	controllers.SetupRoutes(app)

	err := app.Listen(fmt.Sprintf(":%s", config.PORT))
	if err != nil {
		log.Fatalf("couldn't listen to port %s \n error: %s", config.PORT, err.Error())
	}
}
