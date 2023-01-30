package main

import (
	"github.com/kainhuck/irisscaffold/internal/bootstrap"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/config"
	"github.com/kainhuck/irisscaffold/internal/logger"
)

func main() {
	logger.Init(config.Cfg.Logger.LogLevel, config.Cfg.Logger.FileName)

	app := bootstrap.NewApp(config.Cfg.Logger.LogLevel, greet.InitRoutes)
	bootstrap.Run(config.Cfg.Service.Host, config.Cfg.Service.Port, app)
}
