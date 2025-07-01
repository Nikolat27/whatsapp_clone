package data

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
	"time"
)

type ChatModel struct {
	DB *mongo.Database
}

type Chat struct {
	Id           primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Participants []primitive.ObjectID `bson:"participants" json:"participants,omitempty"`
	LastMessage  *Message             `bson:"last_message" json:"last_message,omitempty"`
	CreatedAt    time.Time            `bson:"created_at" json:"created_at,omitempty"`
}

const chatCollection = "chats"

func (c *ChatModel) CreateChatInstance(participants []primitive.ObjectID, lastMessage *Message) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var chat = &Chat{
		Participants: participants,
		LastMessage:  lastMessage,
		CreatedAt:    time.Now(),
	}

	result, err := c.DB.Collection(chatCollection).InsertOne(ctx, chat)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (c *ChatModel) GetChatInstance(chatId primitive.ObjectID) (*Chat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": chatId,
	}

	var chat Chat
	if err := c.DB.Collection(chatCollection).FindOne(ctx, filter).Decode(&chat); err != nil {
		return nil, err
	}

	return &chat, nil
}

func (c *ChatModel) CheckChatExists(participants []primitive.ObjectID) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.M{
		"participants": bson.M{
			"$all": participants,
		},
	}

	var output struct {
		Id primitive.ObjectID `bson:"_id"`
	}

	err := c.DB.Collection(chatCollection).FindOne(ctx, filter).Decode(&output)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	return output.Id.Hex(), err
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
		return errors.New("no chat found matching that ID")
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

	defer func() {
		if err := cursor.Close(ctx); err != nil {
			slog.Error("closing cursor", "error", err)
		}
	}()

	var chats []Chat
	if err = cursor.All(ctx, &chats); err != nil {
		return nil, err
	}

	return chats, nil
}
