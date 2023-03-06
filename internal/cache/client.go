package cache

import (
	"github.com/kainhuck/irisscaffold/internal/cache/redis"
	"github.com/kainhuck/irisscaffold/internal/configx"
)

type Client interface {
}

func NewClient(cfg configx.Redis) Client {
	return redis.NewClient(cfg.Host, cfg.Port, cfg.Password, cfg.DB)
}
