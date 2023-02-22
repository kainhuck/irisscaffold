package application

import (
	"fmt"
	"github.com/kainhuck/irisscaffold/internal/db"
	"github.com/kainhuck/irisscaffold/internal/db/mysql"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/config"
	"os"
)

type Application struct {
	cfg      *config.Config
	dbClient db.Client
}

func NewApplication(cfg *config.Config) *Application {
	dbClient, err := mysql.NewClient(cfg.Database.Mysql.Dsn())
	if err != nil {
		fmt.Printf("new db client failed: %v\n", err)
		os.Exit(-1)
	}
	if err := dbClient.AutoMigrate(); err != nil {
		fmt.Printf("autoMigrate failed: %v\n", err)
		os.Exit(-1)
	}

	return &Application{
		cfg:      cfg,
		dbClient: dbClient,
	}
}