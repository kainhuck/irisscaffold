package controller

import (
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/application"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/config"
)

type Controller struct {
	app *application.Application
}

func NewController(cfg *config.Config) *Controller {
	return &Controller{
		app: application.NewApplication(cfg),
	}
}
