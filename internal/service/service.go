package service

import (
	"taskmaster/internal/models"
	"taskmaster/internal/repository"
)

type Task interface {
	CreateTask(task models.InputTask) (string, int, error)
	GetTasks(status string) ([]models.InputTask, error)
}

type Service struct {
	Task
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Task: newTaskService(repo.Task)}
}
