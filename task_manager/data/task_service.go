package data

import (
	"errors"
	"task_manager/models"
	"github.com/google/uuid"
)

var tasks = []models.Task{}

func GetAllTasks() []models.Task {
	return tasks
}

func GetTaskByID(id string) (*models.Task, error) {
	for i, t := range tasks {
		if t.ID == id {
			return &tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

func CreateTask(task models.Task) models.Task {
	task.ID = uuid.New().String()
	tasks = append(tasks, task)
	return task
}

func UpdateTask(id string, updatedTask models.Task) (*models.Task, error) {
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Title = updatedTask.Title
			tasks[i].Description = updatedTask.Description
			tasks[i].DueDate = updatedTask.DueDate
			tasks[i].Status = updatedTask.Status
			return &tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

func DeleteTask(id string) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}