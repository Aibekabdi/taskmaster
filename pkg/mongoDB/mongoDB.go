package mongodb

import (
	"context"
	"fmt"
	"taskmaster/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(cfg config.MongoDB, ctx context.Context) (*mongo.Client, error) {
	// Установка настройки клиента
	
	clientOpions := options.Client().ApplyURI(fmt.Sprintf("%s://%s:%s", cfg.DBName, cfg.Host, cfg.Port))

	// Подключение к mongodb
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
