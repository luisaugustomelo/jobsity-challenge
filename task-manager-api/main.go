package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	config.LoadEnv()

	// Middleware
	app.Use(logger.New())
	controllers.SetupRoutes(app)

	err := app.Listen(fmt.Sprintf(":%s", config.PORT))
	if err != nil {
		log.Fatalf("couldn't listen to port %s \n error: %s", config.PORT, err.Error())
	}
}
