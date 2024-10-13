package services

import (
	"errors"
	"log"
	"task-manager-api/db"
	"task-manager-api/models"
)

type TaskService interface {
	CreateTask(description string) (*models.Task, error)
	AcceptTask(id uint) error
	DeleteTask(id uint) error
	GetAllTasks() ([]models.Task, error)
}

type taskService struct{}

func NewTaskService() TaskService {
	return &taskService{}
}

func (s *taskService) CreateTask(description string) (*models.Task, error) {
	task := models.Task{
		Description: description,
	}

	result := db.DB.Create(&task)
	if result.Error != nil {
		return nil, result.Error
	}

	log.Printf("Task criada: %v", task)
	return &task, nil
}

func (s *taskService) AcceptTask(id uint) error {
	var task models.Task
	result := db.DB.First(&task, id)
	if result.Error != nil {
		return result.Error
	}

	task.Accept = !task.Accept
	db.DB.Save(&task)

	log.Printf("Task id %d accepted %v", id, task.Accept)
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
