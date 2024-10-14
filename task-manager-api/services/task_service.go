package services

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"task-manager-api/models"
)

type Database interface {
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Save(value interface{}) (tx *gorm.DB)
}

type TaskService interface {
	CreateTask(description, status string) (*models.Task, error)
	UpdateTask(id uint, description, status string) error
	DeleteTask(id uint) error
	GetAllTasks() ([]models.Task, error)
}

type taskService struct {
	db Database
}

func NewTaskService(db Database) TaskService {
	return &taskService{
		db: db,
	}
}

func (s *taskService) CreateTask(description, status string) (*models.Task, error) {
	task := models.Task{
		Description: description,
		Status:      status,
	}

	result := s.db.Create(&task)
	if result.Error != nil {
		return nil, result.Error
	}

	log.Printf("Created task: %v", task)
	return &task, nil
}

func (s *taskService) UpdateTask(id uint, description, status string) error {
	var task models.Task
	result := s.db.First(&task, id)
	if result.Error != nil {
		return result.Error
	}

	task.Description = description
	task.Status = status
	s.db.Save(&task)

	log.Printf("Task id %d updated %v", id, task.Description)
	return nil
}

func (s *taskService) DeleteTask(id uint) error {
	result := s.db.Delete(&models.Task{}, id)
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
	result := s.db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}

	log.Printf("List all tasks: %v", tasks)
	return tasks, nil
}
