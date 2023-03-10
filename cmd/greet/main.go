package main

import (
	"fmt"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	_ "github.com/kainhuck/irisscaffold/docs/greet"
	"github.com/kainhuck/irisscaffold/internal/bootstrap"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/config"
	"github.com/kainhuck/irisscaffold/internal/logger"
)

func main() {
	cfg := config.NewConfig()

	log := logger.NewLogger(cfg.Logger)

	app := bootstrap.NewApp(log)

	greet.InitRoutes(log, app, cfg)

	swagCfg := &swagger.Config{
		URL:         fmt.Sprintf("%s/swagger/doc.json", cfg.Service.BaseURL()),
		DeepLinking: true,
	}

	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(swagCfg, swaggerFiles.Handler))

	bootstrap.Run(log, cfg.Service, app)
}
