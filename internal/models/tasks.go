package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Title     string             `json:"title" bson:"title"`
	ActiveAt  string             `json:"activeAt" bson:"activeAt"`
	Status    string             `json:"status,omitempty" bson:"status"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
}
