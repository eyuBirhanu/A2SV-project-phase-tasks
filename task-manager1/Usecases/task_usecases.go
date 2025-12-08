package usecases

import (
	domain "task-manager/Domain"
	repositories "task-manager/Repositories"
)

type TaskUsecase interface {
	CreateTask(task *domain.Task) error
	GetAllTasks() ([]domain.Task, error)
}

type taskUsecase struct {
	repo repositories.TaskRepository
}

func NewTaskUsecase(repo repositories.TaskRepository) TaskUsecase {
	return &taskUsecase{repo: repo}
}

func (u *taskUsecase) CreateTask(task *domain.Task) error {
	if task.Status == "" {
		task.Status = "Pending"
	}
	return u.repo.Save(task)
}

func (u *taskUsecase) GetAllTasks() ([]domain.Task, error) {
	return u.repo.FetchAll()
}
