package repository

import (
	"context"
	"errors"
	"net/http"
	"taskmaster/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	c *mongo.Collection
}

func newTaskRepository(c *mongo.Collection) *TaskRepository {
	return &TaskRepository{c: c}
}

func (d *TaskRepository) CreateTask(task models.InputTask, activeAt time.Time) (string, int, error) {
	filter := bson.M{"title": task.Title, "activeAt": activeAt}
	err := d.c.FindOne(context.TODO(), filter).Err()
	if err != nil && err != mongo.ErrNoDocuments {
		return "", http.StatusInternalServerError, err
	}
	if err != mongo.ErrNoDocuments {
		return "", http.StatusBadRequest, errors.New("Error occurred while checking task uniqueness")
	}

	insertResult, err := d.c.InsertOne(context.TODO(), bson.M{
		"_id":       primitive.NewObjectID(),
		"title":     task.Title,
		"activeAt":  activeAt,
		"status":    task.Status,
		"createdAt": task.CreatedAt,
	})
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	return insertResult.InsertedID.(primitive.ObjectID).Hex(), 0, nil
}

func (r *TaskRepository) GetTasks() ([]models.Task, error) {
	cur, err := r.c.Find(context.Background(), bson.D{})
	if err != nil {
		// Если произошла ошибка, отправляем ответ со статусом 500
		return nil, err
	}
	defer cur.Close(context.Background())

	// Слайс для хранения задач
	var tasks []models.Task

	// Проходим по всем задачам
	for cur.Next(context.Background()) {
		// Создаем новую задачу
		var task models.Task

		// Декодируем задачу
		err := cur.Decode(&task)
		if err != nil {
			return nil, err
		}

		// Добавляем задачу в слайс
		tasks = append(tasks, task)
	}
	return tasks, nil
}
