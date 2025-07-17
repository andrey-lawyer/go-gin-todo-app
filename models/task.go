package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type StatusType string

const (
	StatusPending    StatusType = "pending"
	StatusInProgress StatusType = "in_progress"
	StatusDone       StatusType = "done"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Status      StatusType         `bson:"status" json:"status"`
	OwnerID     primitive.ObjectID `bson:"owner_id" json:"owner_id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
