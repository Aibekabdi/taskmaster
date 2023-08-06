package mongodb

import (
	"context"
	"fmt"
	"taskmaster/pkg/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(cfg config.MongoDB) (*mongo.Client, error) {
	// Установка настройки клиента
	clientOpions := options.Client().ApplyURI(fmt.Sprintf("%s://%s:%s@%s:%s/?authSource=%s", cfg.DBName, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.AuthSourse))

	// Подключение к mongodb
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOpions)
	if err != nil {
		return nil, err
	}

	// Проверка подключения
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}
