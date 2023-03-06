package application

import (
	"github.com/kainhuck/irisscaffold/internal/cache"
	"github.com/kainhuck/irisscaffold/internal/db"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/config"
	"github.com/kainhuck/irisscaffold/internal/logger"
)

type Application struct {
	log         logger.Logger
	cfg         *config.Config
	dbClient    db.Client
	cacheClient cache.Client
}

func NewApplication(log logger.Logger, cfg *config.Config) *Application {
	return &Application{
		log:         log,
		cfg:         cfg,
		dbClient:    db.NewClient(cfg.Database.Mysql),
		cacheClient: cache.NewClient(cfg.Database.Redis),
	}
}
