package handlers

import (
	"github.com/gofiber/fiber/v2"
	"task-manager-api/services"
)

type TaskHandler struct {
	taskService services.TaskService
}

func NewTaskHandler(taskService services.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	type Request struct {
		Description string `json:"description"`
	}
	req := new(Request)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	task, err := h.taskService.CreateTask(req.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create task"})
	}

	return c.JSON(task)
}

func (h *TaskHandler) AcceptTask(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := h.taskService.AcceptTask(uint(id)); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := h.taskService.DeleteTask(uint(id)); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *TaskHandler) GetAllTasks(c *fiber.Ctx) error {
	tasks, err := h.taskService.GetAllTasks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not list tasks"})
	}
	return c.JSON(tasks)
}
