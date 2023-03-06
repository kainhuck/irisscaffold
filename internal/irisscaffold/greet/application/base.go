package application

import (
	"github.com/kainhuck/irisscaffold/internal/cache"
	"github.com/kainhuck/irisscaffold/internal/cache/redis"
	"github.com/kainhuck/irisscaffold/internal/db"
	"github.com/kainhuck/irisscaffold/internal/db/mysql"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/config"
	"github.com/kainhuck/irisscaffold/internal/logger"
	"os"
)

type Application struct {
	log         logger.Logger
	cfg         *config.Config
	dbClient    db.Client
	cacheClient cache.Client
}

func NewApplication(log logger.Logger, cfg *config.Config) *Application {
	dbClient, err := mysql.NewClient(cfg.Database.Mysql.Dsn())
	if err != nil {
		log.Errorf("new db client failed: %v\n", err)
		os.Exit(-1)
	}
	if err := dbClient.AutoMigrate(); err != nil {
		log.Errorf("autoMigrate failed: %v\n", err)
		os.Exit(-1)
	}

	return &Application{
		log:      log,
		cfg:      cfg,
		dbClient: dbClient,
		cacheClient: redis.NewClient(
			cfg.Database.Redis.Host,
			cfg.Database.Redis.Port,
			cfg.Database.Redis.Password,
			cfg.Database.Redis.DB,
		),
	}
}
