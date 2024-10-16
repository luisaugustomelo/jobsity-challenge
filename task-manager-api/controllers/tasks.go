package controllers

import (
	"github.com/gofiber/fiber/v2"
	"task-manager-api/db"
	"task-manager-api/handlers"
	"task-manager-api/services"
)

func SetupTaskRoutes(app *fiber.App) {
	api := app.Group("/task")

	// Dependency Injection
	taskService := services.NewTaskService(db.DB)
	taskHandler := handlers.NewTaskHandler(taskService)

	registerCreateTaskHandler(api, taskHandler)
}

func registerCreateTaskHandler(api fiber.Router, taskHandler *handlers.TaskHandler) {
	api.Post("/create", taskHandler.CreateTask)
	api.Patch("/:id/update", taskHandler.UpdateTask)
	api.Delete("/:id/delete", taskHandler.DeleteTask)
	api.Get("/allTasks", taskHandler.GetAllTasks)
}
