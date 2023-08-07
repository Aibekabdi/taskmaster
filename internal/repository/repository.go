package repository

import (
	"taskmaster/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Task interface {
	CreateTask(task models.InputTask, activeAt time.Time) (string, int, error)
	GetTasks() ([]models.Task, error)
}

type Repository struct {
	Task
}

func NewRepository(c *mongo.Collection, dbname string) *Repository {
	return &Repository{
		Task: newTaskRepository(c),
	}
}
