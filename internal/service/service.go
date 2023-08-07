package service

import "taskmaster/internal/repository"

type Task interface {
}

type Service struct {
	Task
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Task: newTaskService(repo.Task)}
}
