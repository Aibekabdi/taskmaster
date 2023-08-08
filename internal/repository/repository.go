package repository

import (
	"context"
	"taskmaster/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Task interface {
	CreateTask(ctx context.Context, task models.InputTask, activeAt time.Time, createdAt time.Time) (string, int, error)
	GetTasks(ctx context.Context, status string) ([]models.InputTask, error)
	DeleteTask(ctx context.Context, id string) (int, error)
	MarkTaskDone(ctx context.Context, id string) (int, error)
}

type Repository struct {
	Task
}

func NewRepository(c *mongo.Collection, dbname string) *Repository {
	return &Repository{
		Task: newTaskRepository(c),
	}
}
