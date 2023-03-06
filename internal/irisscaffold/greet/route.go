package greet

import (
	"github.com/kainhuck/irisscaffold/internal/context"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/config"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/controller"
	"github.com/kainhuck/irisscaffold/internal/logger"
	"github.com/kainhuck/irisscaffold/internal/middleware"
	"github.com/kataras/iris/v12"
	"net/http"
)

// @title irisscaffold API
// @version 1.0
// @description iris scaffold for everyone
// @host localhost:8080
// @BasePath /api/v1

func InitRoutes(log logger.Logger, app *iris.Application, cfg *config.Config) {
	v1Route := app.Party("/api/v1")

	ctr := controller.NewController(log, cfg)

	{
		v1Route.Handle(http.MethodGet, "/hello", context.Handler(ctr.GreetHandler))
		v1Route.Handle(http.MethodPost, "/login", context.Handler(ctr.LoginHandler))
	}

	v1Route.Use(middleware.JwtVerify(cfg.Jwt.SigKey))
	{
		v1Route.Handle(http.MethodGet, "/ws", context.Handler(ctr.WebsocketHandler))
		v1Route.Handle(http.MethodGet, "/jwt/demo", context.Handler(ctr.JwtDemoHandler))
		v1Route.Handle(http.MethodPost, "/logout", context.Handler(ctr.LogoutHandler))
	}
}
