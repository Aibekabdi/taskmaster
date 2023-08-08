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
	createdAt := time.Now()
	if activeAt.Before(createdAt) {
		return "", http.StatusBadRequest, errors.New("invalid time")
	}
	return s.taskRepo.CreateTask(task, activeAt, createdAt)
}

func (s *TaskService) GetTasks(status string) ([]models.InputTask, error) {
	if status == "" {
		status = "active"
	}
	return s.taskRepo.GetTasks(status)
}
