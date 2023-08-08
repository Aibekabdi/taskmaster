package service

import (
	"context"
	"taskmaster/internal/models"
	"taskmaster/internal/repository"
)

type Task interface {
	CreateTask(ctx context.Context, task models.InputTask) (string, int, error)
	GetTasks(ctx context.Context, status string) ([]models.InputTask, error)
	DeleteTask(ctx context.Context, id string) (int, error)
	MarkTaskAsDone(ctx context.Context, id string) (int, error)
	UpdateTask(ctx context.Context, updatedInput models.InputTask, id string,) (int, error) 
}

type Service struct {
	Task
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Task: newTaskService(repo.Task)}
}
