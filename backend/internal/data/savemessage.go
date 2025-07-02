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

type SaveMessageModel struct {
	DB *mongo.Database
}

type SaveMessage struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId      primitive.ObjectID `bson:"user_id" json:"user_id"`
	TextContent string             `bson:"text_content"  json:"msg_content"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

const saveMessageCollection = "save_message"

func (m *SaveMessageModel) InsertSaveMessageInstance(userId primitive.ObjectID, payload []byte) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	msg := SaveMessage{
		UserId:      userId,
		TextContent: string(payload),
		CreatedAt:   time.Now(),
	}

	result, err := m.DB.Collection(saveMessageCollection).InsertOne(ctx, msg)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (m *SaveMessageModel) UpdateSaveMessageInstance(msgId primitive.ObjectID, newText []byte) error {
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

	_, err := m.DB.Collection(saveMessageCollection).UpdateOne(ctx, filter, update)
	return err
}

func (m *SaveMessageModel) DeleteSaveMessageInstance(msgId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": msgId,
	}

	_, err := m.DB.Collection(saveMessageCollection).DeleteOne(ctx, filter)
	return err
}

func (m *SaveMessageModel) GetAllSaveMessages(userId primitive.ObjectID) ([]SaveMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.M{
		"user_id": userId,
	}

	findOptions := options.FindOptions{}
	findOptions.SetLimit(10)

	cursor, err := m.DB.Collection(messageCollection).Find(ctx, filter, &findOptions)
	if errors.Is(mongo.ErrNoDocuments, err) {
		return nil, errors.New("no save msg found for this user")
	}

	if err != nil {
		return nil, err
	}

	var Msgs []SaveMessage
	if err = cursor.All(ctx, &Msgs); err != nil {
		return nil, err
	}

	return Msgs, nil
}
