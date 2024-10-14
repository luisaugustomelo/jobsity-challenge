package services

import (
	"errors"
	"log"
	"task-manager-api/db"
	"task-manager-api/models"
)

type TaskService interface {
	CreateTask(description, status string) (*models.Task, error)
	UpdateTask(id uint, description, status string) error
	DeleteTask(id uint) error
	GetAllTasks() ([]models.Task, error)
}

type taskService struct{}

func NewTaskService() TaskService {
	return &taskService{}
}

func (s *taskService) CreateTask(description, status string) (*models.Task, error) {
	task := models.Task{
		Description: description,
		Status:      status,
	}

	result := db.DB.Create(&task)
	if result.Error != nil {
		return nil, result.Error
	}

	log.Printf("Created task: %v", task)
	return &task, nil
}

func (s *taskService) UpdateTask(id uint, description, status string) error {
	var task models.Task
	result := db.DB.First(&task, id)
	if result.Error != nil {
		return result.Error
	}

	task.Description = description
	task.Status = status
	db.DB.Save(&task)

	log.Printf("Task id %d updated %v", id, task.Description)
	return nil
}

func (s *taskService) DeleteTask(id uint) error {
	result := db.DB.Delete(&models.Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("task not found")
	}

	log.Printf("Task id %d deleted", id)
	return nil
}

func (s *taskService) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	result := db.DB.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}

	log.Printf("List all tasks: %v", tasks)
	return tasks, nil
}
