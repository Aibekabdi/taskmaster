package repository

import "go.mongodb.org/mongo-driver/mongo"

type TaskRepository struct {
	c *mongo.Collection
}

func newTaskRepository(c *mongo.Collection) *TaskRepository {
	return &TaskRepository{c: c}
}

