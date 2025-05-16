package data

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Models struct {
	User UserModel
}

func NewModels(db *mongo.Database) *Models {
	return &Models{
		User: UserModel{DB: db},
	}
}
