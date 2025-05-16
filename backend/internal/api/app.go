package api

import (
	"github.com/go-redis/redis"
	"whatsapp_clone/internal/data"
)

type Application struct {
	Models *data.Models
	Rli    *redis.Client
}
