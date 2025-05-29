package data

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ChatModel struct {
	DB *mongo.Database
}

type Chat struct {
	Id           primitive.ObjectID   `bson:"_id,omitempty"`
	Participants []primitive.ObjectID `bson:"participants"`
	LastMassage  *Message             `bson:"last_massage"`
	CreatedAt    time.Time            `bson:"created_at"`
}

var collectionName = "chats"

func (c *ChatModel) CreateChatInstance(participants []primitive.ObjectID, lastMessage *Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	chat := Chat{
		Participants: participants,
		LastMassage:  lastMessage,
		CreatedAt:    time.Now(),
	}

	_, err := c.DB.Collection(collectionName).InsertOne(ctx, chat)
	return err
}

func (c *ChatModel) DeleteChatInstance(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	chat := bson.M{
		"_id": id,
	}

	res, err := c.DB.Collection(collectionName).DeleteOne(ctx, chat)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no chat found with that id")
	}
	return nil
}
