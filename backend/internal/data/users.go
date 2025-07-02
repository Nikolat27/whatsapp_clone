package data

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	passwordHelper "whatsapp_clone/internal/helpers/password"
)

type UserModel struct {
	DB *mongo.Database
}

type User struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	Username      string             `bson:"username"`
	Password      string             `bson:"password"`
	Name          string             `bson:"name"`
	About         string             `bson:"about"`
	ProfilePicUrl string             `bson:"profile_url"`
	CreatedAt     time.Time          `bson:"createdAt"`
}

const userCollection = "users"

func (u *UserModel) CreateUserInstance(username, password string) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	hashedPassword, err := passwordHelper.HashPassword(password)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to hash password %s", err)
	}

	user := User{
		Username:  username,
		Password:  hashedPassword,
		CreatedAt: time.Now().UTC(),
	}

	result, err := u.DB.Collection(userCollection).InsertOne(ctx, user)
	return result.InsertedID.(primitive.ObjectID), nil
}

func (u *UserModel) GetUserInstanceByUsername(username string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.M{
		"username": username,
	}

	var user User
	err := u.DB.Collection(userCollection).FindOne(ctx, filter).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to find user '%s': %w", username, err)
	}

	return &user, nil
}

func (u *UserModel) GetUserInstanceById(id primitive.ObjectID) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": id,
	}

	var user User
	err := u.DB.Collection(userCollection).FindOne(ctx, filter).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to find user '%s': %w", id, err)
	}

	return &user, nil
}

func (u *UserModel) UpdateUserInstance(userId primitive.ObjectID, updates bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": userId,
	}

	update := bson.M{
		"$set": updates,
	}

	_, err := u.DB.Collection(userCollection).UpdateOne(ctx, filter, update)
	return err
}

func (u *UserModel) DeleteUserInstance(userId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": userId,
	}

	result, err := u.DB.Collection(userCollection).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no user found to delete. 404")
	}

	return nil
}

func (u *UserModel) SearchUser(username string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.M{
		"username": username,
	}

	_ = u.DB.Collection(userCollection).FindOne(ctx, filter)
}
