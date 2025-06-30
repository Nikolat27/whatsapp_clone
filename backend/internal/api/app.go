package api

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"
	"whatsapp_clone/cipher"
	"whatsapp_clone/internal/data"
)

type Application struct {
	Server          *http.Server
	Models          *data.Models
	Rli             *redis.Client
	ChatConnections map[string]map[string]*websocket.Conn
	ConnMu          sync.Mutex
	PublicKeys      map[string]string // {"userid": "public key"}
	Cipher          *cipher.Cipher
}

func NewApp() (*Application, error) {
	client, err := setupMongoDB()
	if err != nil {
		return nil, err
	}

	// Connecting mongo
	dbName := getDatabaseName()
	db := client.Database(dbName)

	// Connecting redis
	rli, err := setupRedis()
	if err != nil {
		return nil, err
	}

	cip, err := cipher.InitCipher()
	if err != nil {
		return nil, err
	}

	// Creating the models objects
	models := data.NewModels(db)
	app := &Application{
		Models:          models,
		Rli:             rli,
		Cipher:          cip,
		ChatConnections: make(map[string]map[string]*websocket.Conn),
	}

	srv := app.setupHttpServer()
	app.Server = srv
	
	return app, nil
}

func (app *Application) setupHttpServer() *http.Server {
	port := getPortHTTP()

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: routes(app),
	}
}

func (app *Application) RunServer() error {
	return app.Server.ListenAndServe()
}

func setupMongoDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dsn := getMongoDBDsn()

	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	slog.Info("mongodb", "result", "connected successfully")
	return client, nil
}

func setupRedis() (*redis.Client, error) {
	dsn := getRedisDsn()
	client := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func getMongoDBDsn() string {
	dsn := os.Getenv("MONGODB_DSN")
	if dsn == "" {
		panic("mongodb dsn env variable doesn't exist")
	}
	return dsn
}

func getDatabaseName() string {
	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		panic("database name env variable does not exist")
	}
	return dbName
}

func getRedisDsn() string {
	dsn := os.Getenv("REDIS_DSN")
	if dsn == "" {
		panic("redis_dsn env variable does not exist")
	}

	return dsn
}

func getPortHTTP() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8000" // default port
	}
	return port
}
