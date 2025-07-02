package data

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Models struct {
	User        UserModel
	Chat        ChatModel
	Message     MessageModel
	SaveMessage SaveMessageModel
}

func NewModels(db *mongo.Database) *Models {
	return &Models{
		User:        UserModel{DB: db},
		Chat:        ChatModel{DB: db},
		Message:     MessageModel{DB: db},
		SaveMessage: SaveMessageModel{DB: db},
	}
}
