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
	Id           primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Participants []primitive.ObjectID `bson:"participants" json:"participants,omitempty"`
	LastMassage  *Message             `bson:"last_massage" json:"last_message,omitempty"`
	CreatedAt    time.Time            `bson:"created_at" json:"created_at,omitempty"`
}

const chatCollection = "chats"

func (c *ChatModel) CreateChatInstance(participants []primitive.ObjectID, lastMessage *Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	chat := Chat{
		Participants: participants,
		LastMassage:  lastMessage,
		CreatedAt:    time.Now(),
	}

	_, err := c.DB.Collection(chatCollection).InsertOne(ctx, chat)
	return err
}

func (c *ChatModel) GetChatInstance(chatId primitive.ObjectID) (*Chat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": chatId,
	}

	var chat Chat
	err := c.DB.Collection(chatCollection).FindOne(ctx, filter).Decode(&chat)
	if err != nil {
		return nil, err
	}
	return &chat, nil
}

func (c *ChatModel) DeleteChatInstance(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	chat := bson.M{
		"_id": id,
	}

	res, err := c.DB.Collection(chatCollection).DeleteOne(ctx, chat)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no chat found with that id")
	}
	return nil
}

func (c *ChatModel) GetUserChats(userId primitive.ObjectID) ([]Chat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.M{
		"participants": userId,
	}

	cursor, err := c.DB.Collection(chatCollection).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var chats []Chat
	err = cursor.All(ctx, &chats)
	if err != nil {
		return nil, err
	}

	return chats, nil
}
