package main

import (
	"github.com/kainhuck/irisscaffold/internal/bootstrap"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/config"
	"github.com/kainhuck/irisscaffold/internal/logger"
)

func main() {
	cfg := config.NewConfig()

	logger.Init(cfg.Logger)

	app := bootstrap.NewApp(cfg.Logger.LogLevel)

	greet.InitRoutes(app, cfg)

	bootstrap.Run(cfg.Service, app)
}
