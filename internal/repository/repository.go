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
	MarkTaskAsDone(ctx context.Context, id string) (int, error)
	UpdateTask(ctx context.Context, updatedInput models.InputTask, id string, activeAt time.Time) (int, error) 
}

type Repository struct {
	Task
}

func NewRepository(c *mongo.Collection, dbname string) *Repository {
	return &Repository{
		Task: newTaskRepository(c),
	}
}
