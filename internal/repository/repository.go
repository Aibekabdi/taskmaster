package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
}

func NewRepository(db *mongo.Client, dbname string) *Repository {
	return &Repository{}
}
