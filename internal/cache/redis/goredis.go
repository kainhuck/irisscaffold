package redis

import (
	"fmt"
	"github.com/kainhuck/irisscaffold/internal/cache"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	rdb *redis.Client
}

func NewClient(host string, port int, password string, db int) cache.Client {

	return &Client{
		rdb: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", host, port),
			Password: password,
			DB:       db,
		}),
	}
}
