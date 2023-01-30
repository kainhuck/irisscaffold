package main

import (
	"github.com/kainhuck/irisscaffold/internal/bootstrap"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/config"
	"github.com/kainhuck/irisscaffold/internal/logger"
)

func init() {
	logger.Init(config.Cfg.Logger)
}

func main() {
	app := bootstrap.NewApp(config.Cfg.Logger, greet.InitRoutes)
	bootstrap.Run(config.Cfg.Service, app)
}
