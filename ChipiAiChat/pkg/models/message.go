package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Message struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	RoomID        string             `bson:"roomId"`
	Username      string             `bson:"username"`
	Content       string             `bson:"content"`
	ContentWithAi string             `bson:"isAi"`
	CreatedAt     time.Time          `bson:"createdAt"`
}
