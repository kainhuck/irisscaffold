package application

import (
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/config"
)

type Application struct {
	cfg *config.Config
}

func NewApplication(cfg *config.Config) *Application {

	return &Application{
		cfg: cfg,
	}
}
