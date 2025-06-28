package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MessageModel struct {
	DB *mongo.Database
}

type Message struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	ChatId      primitive.ObjectID `bson:"chat_id"`
	SenderId    primitive.ObjectID `bson:"sender_id"`
	TextContent string             `bson:"text_content"`
	CreatedAt   time.Time          `bson:"created_at"`
}

const messageCollection = "messages"

func (m *MessageModel) InsertMessageInstance(chatId, senderId primitive.ObjectID, textContent []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	msg := Message{
		ChatId:      chatId,
		SenderId:    senderId,
		TextContent: string(textContent),
		CreatedAt:   time.Now(),
	}

	_, err := m.DB.Collection(messageCollection).InsertOne(ctx, msg)
	return err
}

func (m *MessageModel) UpdateMessageInstance(msgId primitive.ObjectID, newText []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": msgId,
	}

	update := bson.M{
		"$set": bson.M{
			"text_content": string(newText),
		},
	}

	_, err := m.DB.Collection(messageCollection).UpdateOne(ctx, filter, update)
	return err
}

func (m *MessageModel) DeleteMessageInstance(msgId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": msgId,
	}

	_, err := m.DB.Collection(messageCollection).DeleteOne(ctx, filter)
	return err
}
