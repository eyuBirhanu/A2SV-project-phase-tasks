package usecases

import (
	"errors"
	domain "task-manager/Domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepo struct {
	mock.Mock
}

func (m *MockTaskRepo) Save(task *domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepo) FetchAll() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *MockTaskRepo) FetchByID(id string) (*domain.Task, error) { return nil, nil }
func (m *MockTaskRepo) Delete(id string) error                    { return nil }

func TestCreateTask_Success(t *testing.T) {
	mockRepo := new(MockTaskRepo)
	usecase := NewTaskUsecase(mockRepo)

	newTask := &domain.Task{
		ID:    "1",
		Title: "Test Task",
	}

	mockRepo.On("Save", newTask).Return(nil)

	err := usecase.CreateTask(newTask)

	assert.NoError(t, err)
	assert.Equal(t, "Pending", newTask.Status)
	mockRepo.AssertExpectations(t)
}

func TestFetchAll_Error(t *testing.T) {
	mockRepo := new(MockTaskRepo)
	usecase := NewTaskUsecase(mockRepo)

	mockRepo.On("FetchAll").Return([]domain.Task{}, errors.New("database down"))

	tasks, err := usecase.GetAllTasks()

	assert.Error(t, err)
	assert.Empty(t, tasks)
	assert.Equal(t, "database down", err.Error())
	mockRepo.AssertExpectations(t)
}
