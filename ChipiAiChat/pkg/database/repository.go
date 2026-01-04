package database

import (
	"ChipiAiChat/pkg/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func SaveMessage(roomId, username, content string) error {
	msg := models.Message{
		RoomID:    roomId,
		Username:  username,
		Content:   content,
		CreatedAt: time.Now(),
	}

	_, err := Messages().InsertOne(context.Background(), msg)
	if err != nil {
		return err
	}

	return nil
}

func CreateNewRoom(nameRoom string) (primitive.ObjectID, error) {
	room := models.Room{
		RoomId:    primitive.NewObjectID(),
		Name:      nameRoom,
		CreatedAt: time.Now(),
	}

	_, err := Rooms().InsertOne(context.Background(), room)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return room.RoomId, nil
}

func UpdateMessageContent(roomId string, username string, newContent string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"roomId":   roomId,
		"username": username,
	}

	update := bson.M{
		"$set": bson.M{
			"isAi":      newContent,
			"updatedAt": time.Now(),
		},
	}

	_, err := Messages().UpdateOne(ctx, filter, update)
	return err
}
