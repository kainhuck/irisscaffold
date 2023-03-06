package controller

import (
	"github.com/gorilla/websocket"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/application"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/config"
	"github.com/kainhuck/irisscaffold/internal/logger"
)

type Controller struct {
	log      logger.Logger
	app      *application.Application
	upGrader websocket.Upgrader
}

func NewController(log logger.Logger, cfg *config.Config) *Controller {
	return &Controller{
		log: log,
		app: application.NewApplication(log, cfg),
		upGrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}
