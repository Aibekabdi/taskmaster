package repository

import (
	"context"
	"errors"
	"net/http"
	"sort"
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

func (d *TaskRepository) CreateTask(task models.InputTask, activeAt time.Time, createdAt time.Time) (string, int, error) {
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
		"status":    "active",
		"createdAt": createdAt,
	})
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	return insertResult.InsertedID.(primitive.ObjectID).Hex(), 0, nil
}

func (r *TaskRepository) GetTasks(status string) ([]models.InputTask, error) {
	// Фильтруем задачи по статусу и активной дате
	filter := bson.M{"status": status}
	if status == "active" {
		filter["activeAt"] = bson.M{"$lte": time.Now()}
	}

	cur, err := r.c.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var tasks []models.Task

	// Проходим по всем задачам
	for cur.Next(context.Background()) {
		var task models.Task

		err := cur.Decode(&task)
		if err != nil {
			return nil, err
		}

		// Проверяем, является ли день активации выходным
		if task.ActiveAt.Weekday() == time.Saturday || task.ActiveAt.Weekday() == time.Sunday {
			task.Title = "ВЫХОДНОЙ - " + task.Title
		}

		tasks = append(tasks, task)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Сортируем задачи по дате создания
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].CreatedAt.Before(tasks[j].CreatedAt)
	})
	var res []models.InputTask
	for _, tasks := range tasks {
		res = append(res, models.InputTask{
			Title:    tasks.Title,
			ActiveAt: tasks.ActiveAt.Format("2006-01-02"),
		})
	}
	return res, nil
}
