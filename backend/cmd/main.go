package main

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"net/http"
	"whatsapp_clone/internal/api"
	"whatsapp_clone/internal/data"
	"whatsapp_clone/internal/routes"
)

func main() {
	client, err := openDB("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	defer func() {
		err = client.Disconnect(context.Background())
		if err != nil {
			panic(err)
		}
	}()

	// Connecting mongo
	db := client.Database("whatsapp_clone")
	if err != nil {
		panic(err)
	}

	// Connecting redis
	rli, err := establishRedis("localhost:6379")
	if err != nil {
		panic(err)
	}

	// Creating the models objects
	models := data.NewModels(db)

	app := &api.Application{
		Models:          models,
		Rli:             rli,
		ChatConnections: make(map[string]map[string]*websocket.Conn),
	}

	server := &http.Server{Addr: ":5000", Handler: routes.Route(app)}

	// Starting the http server
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func openDB(url string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(url)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	slog.Info("mongodb", "result", "connected successfully")

	return client, nil
}

func establishRedis(url string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
