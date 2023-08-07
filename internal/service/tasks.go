package service

import (
	"errors"
	"net/http"
	"taskmaster/internal/models"
	"taskmaster/internal/repository"
	"time"
)

type TaskService struct {
	taskRepo repository.Task
}

func newTaskService(taskRepo repository.Task) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) CreateTask(task models.InputTask) (string, int, error) {
	activeAt, err := time.Parse("2006-01-02", task.ActiveAt)
	if err != nil {
		return "", http.StatusBadRequest, err
	}
	task.CreatedAt = time.Now()
	if activeAt.Before(task.CreatedAt) {
		return "", http.StatusBadRequest, errors.New("invalid time")
	}
	task.Status = "active"
	return s.taskRepo.CreateTask(task, activeAt)
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
	return s.taskRepo.GetTasks()
}
