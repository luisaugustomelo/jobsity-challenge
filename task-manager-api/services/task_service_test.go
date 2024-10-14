package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"task-manager-api/models"
	"testing"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return &gorm.DB{Error: args.Error(0)}
}

func (m *MockDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	if args.Get(0) != nil {
		return args.Get(0).(*gorm.DB)
	}
	return &gorm.DB{Error: nil}
}

func (m *MockDB) Save(value interface{}) *gorm.DB {
	args := m.Called(value)
	if args.Get(0) != nil {
		return args.Get(0).(*gorm.DB)
	}
	return &gorm.DB{Error: nil}
}

func (m *MockDB) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(value, conds)
	if args.Get(0) != nil {
		return args.Get(0).(*gorm.DB)
	}
	return &gorm.DB{Error: nil, RowsAffected: args.Get(1).(int64)}
}

func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	if args.Get(0) != nil {
		return args.Get(0).(*gorm.DB)
	}
	return &gorm.DB{Error: nil}
}

func TestCreateTask(t *testing.T) {
	mockDB := new(MockDB)

	taskService := NewTaskService(mockDB)

	task := &models.Task{
		Description: "Test Task",
		Status:      "doing",
	}

	mockDB.On("Create", task).Return(nil)

	createdTask, err := taskService.CreateTask(task.Description, task.Status)

	assert.Nil(t, err)
	assert.NotNil(t, createdTask)
	assert.Equal(t, "Test Task", createdTask.Description)
	assert.Equal(t, "doing", createdTask.Status)

	mockDB.AssertCalled(t, "Create", task)
}

func TestUpdateTask(t *testing.T) {
	mockDB := new(MockDB)
	taskService := NewTaskService(mockDB)

	task := models.Task{
		ID:          1,
		Description: "Updated Task",
		Status:      "completed",
	}

	mockDB.On("First", mock.AnythingOfType("*models.Task"), []interface{}{task.ID}).Return(&gorm.DB{Error: nil})
	mockDB.On("Save", mock.AnythingOfType("*models.Task")).Return(&gorm.DB{Error: nil})

	err := taskService.UpdateTask(task.ID, "Updated Task", "completed")

	assert.Nil(t, err)
	assert.Equal(t, "Updated Task", task.Description)
	assert.Equal(t, "completed", task.Status)

	mockDB.AssertCalled(t, "First", mock.AnythingOfType("*models.Task"), []interface{}{uint(1)})
}

func TestDeleteTask(t *testing.T) {
	mockDB := new(MockDB)
	taskService := NewTaskService(mockDB)

	task := models.Task{
		ID: 1,
	}

	mockDB.On("Delete", mock.AnythingOfType("*models.Task"), []interface{}{task.ID}).Return(&gorm.DB{Error: nil, RowsAffected: 1})

	err := taskService.DeleteTask(task.ID)

	assert.Nil(t, err)
}

func TestGetAllTasks(t *testing.T) {
	mockDB := new(MockDB)
	taskService := NewTaskService(mockDB)

	tasks := []models.Task{
		{ID: 1, Description: "Test Task 1", Status: "doing"},
		{ID: 2, Description: "Test Task 2", Status: "completed"},
	}

	mockDB.On("Find", mock.AnythingOfType("*[]models.Task"), mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*[]models.Task)
		*arg = tasks
	}).Return(&gorm.DB{Error: nil})

	allTasks, err := taskService.GetAllTasks()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(allTasks))
	assert.Equal(t, "Test Task 1", allTasks[0].Description)
	assert.Equal(t, "Test Task 2", allTasks[1].Description)

	mockDB.AssertCalled(t, "Find", mock.AnythingOfType("*[]models.Task"), mock.Anything)
}
