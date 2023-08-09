compose-build:
	docker-compose build

swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g ./cmd/main.go

compose-up:
	docker-compose up
