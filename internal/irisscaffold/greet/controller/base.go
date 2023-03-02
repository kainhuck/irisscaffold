package controller

import (
	"github.com/gorilla/websocket"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/application"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/config"
)

type Controller struct {
	app      *application.Application
	upGrader websocket.Upgrader
}

func NewController(cfg *config.Config) *Controller {
	return &Controller{
		app: application.NewApplication(cfg),
		upGrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}
