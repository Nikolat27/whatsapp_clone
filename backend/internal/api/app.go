package api

import (
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"sync"
	"whatsapp_clone/internal/data"
)

type Application struct {
	Models          *data.Models
	Rli             *redis.Client
	ChatConnections map[string]map[string]*websocket.Conn
	ConnMu          sync.Mutex
}
