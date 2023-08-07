package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"taskmaster/internal/http"
	"taskmaster/internal/repository"
	"taskmaster/internal/server"
	"taskmaster/internal/service"
	"taskmaster/pkg/config"
	mongodb "taskmaster/pkg/mongoDB"
	"time"
)

func main() {
	// Основа микросервиса

	// Загрузка конфигурации
	conf, err := config.NewConfig("./configs/config.json")
	if err != nil {
		log.Fatalf("failed to load configs: %s", err)
	}

	ctxForDB := context.TODO()
	// подключение в mongodb
	db, err := mongodb.NewMongoDB(conf.MongoDB, ctxForDB)
	if err != nil {
		log.Fatalf("failed to initialize mongo db: %s", err.Error())
	}

	defer func() {
		if err = db.Disconnect(ctxForDB); err != nil {
			log.Fatal("can't close connection db, err:", err)
		} else {
			log.Println("db closed")
		}
	}()

	// Подготовка слоенную архитектуру
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := http.NewHandler(service)

	// Запуск сервера
	srv := new(server.Server)
	go func() {
		if err := srv.Run(conf.Api.Port, handler.InitRoutes()); err != nil {
			log.Printf("error occured while running http server: %s", err.Error())
			return
		}
	}()

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	fmt.Println()
	log.Println("Received terminate, graceful shutdown", sig)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	if err = srv.Stop(ctx); err != nil {
		log.Fatal(err)
	}
}
