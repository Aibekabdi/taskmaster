package repository

import "go.mongodb.org/mongo-driver/mongo"

type Task interface {
}

type Repository struct {
	Task
}

func NewRepository(c *mongo.Collection, dbname string) *Repository {
	return &Repository{
		Task: newTaskRepository(c),
	}
}
