package repositories

import (
	"errors"
	domain "task-manager/Domain"
)

type TaskRepository interface {
	Save(task *domain.Task) error
	FetchAll() ([]domain.Task, error)
	FetchByID(id string) (*domain.Task, error)
	Delete(id string) error
}

type InMemoryTaskRepo struct {
	tasks []domain.Task
}

func NewTaskRepository() TaskRepository {
	return &InMemoryTaskRepo{tasks: []domain.Task{}}
}

func (r *InMemoryTaskRepo) Save(task *domain.Task) error {
	r.tasks = append(r.tasks, *task)
	return nil
}

func (r *InMemoryTaskRepo) FetchAll() ([]domain.Task, error) {
	return r.tasks, nil
}

func (r *InMemoryTaskRepo) FetchByID(id string) (*domain.Task, error) {
	for _, t := range r.tasks {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errors.New("task not found")
}

func (r *InMemoryTaskRepo) Delete(id string) error {
	for i, t := range r.tasks {
		if t.ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
