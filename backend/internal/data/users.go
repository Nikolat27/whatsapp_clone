package data

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"whatsapp_clone/internal/helpers"
)

type UserModel struct {
	DB *mongo.Database
}

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"createdAt"`
}

const userCollection = "users"

func (u *UserModel) CreateUserInstance(username, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password %s", err)
	}

	user := User{
		Username:  username,
		Password:  hashedPassword,
		CreatedAt: time.Now().UTC(),
	}

	_, err = u.DB.Collection(userCollection).InsertOne(ctx, user)
	return err
}

func (u *UserModel) GetUserInstance(username string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.M{
		"username": username,
	}

	var user User
	err := u.DB.Collection(userCollection).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find user '%s': %w", username, err)
	}
	return &user, nil
}

func (u *UserModel) DeleteUserInstance(userId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": userId,
	}
	
	result,  err := u.DB.Collection(userCollection).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no user found to delete. 404")
	}

	return nil
}
