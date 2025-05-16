package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"net/http"
	"whatsapp_clone/internal/api"
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
	
	database := client.Database("whatsapp_clone")
	err = database.CreateCollection(context.Background(), "users")	
	if err != nil {
		panic(err)
	}
	
	app := &api.Application{}
	server := &http.Server{Addr: ":8000", Handler: routes.Route(app)}

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

	slog.Info("mongodb", "connected successfully")

	return client, nil
}


