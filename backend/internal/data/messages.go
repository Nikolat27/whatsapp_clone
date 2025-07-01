package data

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MessageModel struct {
	DB *mongo.Database
}

type Message struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ChatId      primitive.ObjectID `bson:"chat_id" json:"chat_id"`
	SenderId    primitive.ObjectID `bson:"sender_id" json:"sender_id"`
	TextContent string             `bson:"text_content"  json:"msg_content"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

const messageCollection = "messages"

func (m *MessageModel) InsertMessageInstance(chatId, senderId primitive.ObjectID, payload []byte) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	msg := Message{
		ChatId:      chatId,
		SenderId:    senderId,
		TextContent: string(payload),
		CreatedAt:   time.Now(),
	}

	result, err := m.DB.Collection(messageCollection).InsertOne(ctx, msg)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
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

func (m *MessageModel) GetAllMessages(chatId primitive.ObjectID) ([]Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.M{
		"chat_id": chatId,
	}

	findOptions := options.FindOptions{}
	findOptions.SetLimit(10)

	cursor, err := m.DB.Collection(messageCollection).Find(ctx, filter, &findOptions)
	if errors.Is(mongo.ErrNoDocuments, err) {
		return nil, errors.New("no msg found for this chat")
	}

	if err != nil {
		return nil, err
	}

	var Msgs []Message
	if err = cursor.All(ctx, &Msgs); err != nil {
		return nil, err
	}

	return Msgs, nil
}
