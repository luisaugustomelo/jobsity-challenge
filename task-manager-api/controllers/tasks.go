package controllers

import (
	"github.com/gofiber/fiber/v2"
	"task-manager-api/handlers"
	"task-manager-api/services"
)

func SetupTaskRoutes(app *fiber.App) {
	api := app.Group("/task")

	// Dependency Injection
	taskService := services.NewTaskService()
	taskHandler := handlers.NewTaskHandler(taskService)

	registerCreateTaskHandler(api, taskHandler)
}

func registerCreateTaskHandler(api fiber.Router, taskHandler *handlers.TaskHandler) {
	api.Post("/create", taskHandler.CreateTask)
	api.Patch("/:id/accept", taskHandler.AcceptTask)
	api.Delete("/:id/delete", taskHandler.DeleteTask)
	api.Get("/allTasks", taskHandler.GetAllTasks)
}
