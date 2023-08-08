package service

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"taskmaster/internal/models"
	"taskmaster/internal/repository"
	"time"
	"unicode"
)

type TaskService struct {
	taskRepo repository.Task
}

func newTaskService(taskRepo repository.Task) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) CreateTask(ctx context.Context, task models.InputTask) (string, int, error) {
	activeAt, err := time.Parse("2006-01-02", task.ActiveAt)
	if err != nil {
		return "", http.StatusBadRequest, err
	}
	createdAt := time.Now()
	if activeAt.Before(createdAt) {
		return "", http.StatusBadRequest, errors.New("invalid time")
	}
	return s.taskRepo.CreateTask(ctx, task, activeAt, createdAt)
}

func (s *TaskService) GetTasks(ctx context.Context, status string) ([]models.InputTask, error) {
	text := strings.TrimFunc(status, func(r rune) bool {
		return unicode.IsSpace(r)
	})
	if text == "" {
		status = "active"
	}
	return s.taskRepo.GetTasks(ctx, status)
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) (int, error) {
	text := strings.TrimFunc(id, func(r rune) bool {
		return unicode.IsSpace(r)
	})
	if text == "" {
		return http.StatusBadRequest, errors.New("invalid id")
	}
	return s.taskRepo.DeleteTask(ctx, id)
}

func (s *TaskService) MarkTaskDone(ctx context.Context, id string) (int, error) {
	text := strings.TrimFunc(id, func(r rune) bool {
		return unicode.IsSpace(r)
	})
	if text == "" {
		return http.StatusBadRequest, errors.New("invalid id")
	}
	return s.taskRepo.MarkTaskDone(ctx, id)
}
