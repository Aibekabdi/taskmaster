package service

import "taskmaster/internal/repository"

type Service struct {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}