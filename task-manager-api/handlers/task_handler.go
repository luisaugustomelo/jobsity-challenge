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
		Status      string `json:"status"`
	}
	req := new(Request)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	task, err := h.taskService.CreateTask(req.Description, req.Status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create task"})
	}

	return c.JSON(task)
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	type Request struct {
		Description string `json:"description"`
		Status      string `json:"status"`
	}
	req := new(Request)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	id, _ := c.ParamsInt("id")

	if err := h.taskService.UpdateTask(uint(id), req.Description, req.Status); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
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
