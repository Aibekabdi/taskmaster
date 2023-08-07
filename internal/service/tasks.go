package service

import "taskmaster/internal/repository"

type TaskService struct {
	taskRepo repository.Task
}

func newTaskService(taskRepo repository.Task) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}
